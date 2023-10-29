package repository

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rochaeduardo997/money-back/internal/Debtod/entity"
)

type DebtodRepositoryDB struct {
	db *sql.DB
}

func NewDebtodRepository(db *sql.DB) (debtodRepository *DebtodRepositoryDB) {
	debtodRepository = &DebtodRepositoryDB{db: db}
	return
}

func (dr *DebtodRepositoryDB) Save(d *entity.Debtod) (*entity.Debtod, error) {
	err := d.Validate()
	if err != nil {
		return nil, err
	}

	d.Id = uuid.NewString()

	tx, err := dr.db.Begin()
	if err != nil {
		return nil, err
	}

	var debtod entity.Debtod
	err = InsertDebtod(tx, d, &debtod)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = InsertContacts(tx, d, &debtod)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(d.Addresses) > 0 {
		err := InsertAddresses(tx, d, &debtod)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = InsertInvoices(tx, d, &debtod)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &debtod, nil
}

func (dr *DebtodRepositoryDB) GetAll() ([]*entity.Debtod, error) {
	var result []*entity.Debtod

	debtodSQL := "SELECT * FROM tbl_debtods"
	debtodRows, err := dr.db.Query(debtodSQL)
	if err != nil {
		return nil, err
	}

	for debtodRows.Next() {
		var debtod entity.Debtod

		err = ScanDebtods(debtodRows, &debtod)
		if err != nil {
			return nil, err
		}
		err = ScanDebtodContacts(dr.db, &debtod)
		if err != nil {
			return nil, err
		}
		err = ScanDebtodAddresses(dr.db, &debtod)
		if err != nil {
			return nil, err
		}
		err = ScanDebtodInvoices(dr.db, &debtod)
		if err != nil {
			return nil, err
		}

		result = append(result, &debtod)
	}
	return result, nil
}

func (dr *DebtodRepositoryDB) GetBy(id string) (*entity.Debtod, error) {
	debtodSQL := "SELECT * FROM tbl_debtods WHERE id = $1"
	debtodRow := dr.db.QueryRow(debtodSQL, &id)

	var d entity.Debtod
	var personType string
	err := debtodRow.Scan(&d.Id, &d.Name, &d.Surname, &d.BussinessName, &d.Observation, &d.CPF_CNPJ, &personType)
	if err != nil {
		return nil, err
	}

	DeterminePersonType(&d, &personType)

	err = ScanDebtodAddresses(dr.db, &d)
	if err != nil {
		return nil, err
	}

	err = ScanDebtodContacts(dr.db, &d)
	if err != nil {
		return nil, err
	}

	err = ScanDebtodInvoices(dr.db, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (dr *DebtodRepositoryDB) CloseDB() {
	dr.db.Close()
}

func InsertDebtod(tx *sql.Tx, d *entity.Debtod, debtod *entity.Debtod) (err error) {
	debtodSQL := `
		INSERT INTO tbl_debtods(id, name, surname, bussiness_name, observation, cpf_cnpj, person_type)
		VALUES($1,$2,$3,$4,$5,$6,$7)
		RETURNING *
	`
	var personType string
	err = tx.QueryRow(debtodSQL, &d.Id, &d.Name, &d.Surname, &d.BussinessName, &d.Observation, &d.CPF_CNPJ, GetPersonTypeString(&d.PersonType)).Scan(&debtod.Id, &debtod.Name, &debtod.Surname, &debtod.BussinessName, &debtod.Observation, &debtod.CPF_CNPJ, &personType)
	switch strings.ToUpper(personType) {
	case "PF":
		debtod.PersonType = 1
	case "PJ":
		debtod.PersonType = 2
	case "FOREIGN":
		debtod.PersonType = 3
	}

	return
}

func GetPersonTypeString(personType *entity.EPersonType) string {
	switch *personType {
	case 1:
		return "pf"
	case 2:
		return "pj"
	case 3:
		return "foreign"
	default:
		return "pf"
	}
}

func InsertContacts(tx *sql.Tx, d *entity.Debtod, debtod *entity.Debtod) (err error) {
	numberSQL := `
	  INSERT INTO tbl_debtods_contact_numbers(number, fk_debtod_id)
		VALUES($1,$2)
		RETURNING number
	`
	for i := 0; i < len(d.Contacts.Numbers); i++ {
		var number int
		err = tx.QueryRow(numberSQL, &d.Contacts.Numbers[i], &d.Id).Scan(&number)
		if err != nil {
			return
		}
		debtod.Contacts.Numbers = append(debtod.Contacts.Numbers, number)
	}

	emailSQL := `
	  INSERT INTO tbl_debtods_contact_emails(email, fk_debtod_id)
		VALUES($1,$2)
		RETURNING email
	`
	for i := 0; i < len(d.Contacts.Numbers); i++ {
		var email string
		err = tx.QueryRow(emailSQL, &d.Contacts.Emails[i], &d.Id).Scan(&email)
		if err != nil {
			return
		}
		debtod.Contacts.Emails = append(debtod.Contacts.Emails, email)
	}
	return
}

func InsertAddresses(tx *sql.Tx, d *entity.Debtod, debtod *entity.Debtod) (err error) {
	addressSQL := `
		INSERT INTO tbl_debtods_addresses
		VALUES($1,$2,$3,$4,$5,$6,$7)
		RETURNING *
	`
	for i := 0; i < len(d.Addresses); i++ {
		a := &d.Addresses[i]
		var address entity.Address
		err = tx.QueryRow(addressSQL, &a.Street, &a.Number, &a.Zipcode, &a.Neighborhood, &a.Observation, &a.Description, &d.Id).Scan(&address.Street, &address.Number, &address.Zipcode, &address.Neighborhood, &address.Observation, &address.Description, &d.Id)
		if err != nil {
			return
		}
		debtod.Addresses = append(debtod.Addresses, address)
	}

	return
}

func InsertInvoices(tx *sql.Tx, d *entity.Debtod, debtod *entity.Debtod) (err error) {
	invoiceSQL := `
		INSERT INTO tbl_debtods_invoices
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING *
	`
	for i := 0; i < len(d.Addresses); i++ {
		i := &d.Invoices[i]
		i.Id = uuid.NewString()
		var invoice entity.Invoice
		err = tx.QueryRow(invoiceSQL, &i.Id, &i.Identification, &i.Status, &i.DelayNotified, &i.DontNotifyUntil, &i.Observation, &i.Description, i.ExpiredDate.Format(time.RFC3339), &d.Id).Scan(&invoice.Id, &invoice.Identification, &invoice.Status, &invoice.DelayNotified, &invoice.DontNotifyUntil, &invoice.Observation, &invoice.Description, &invoice.ExpiredDate, &d.Id)
		if err != nil {
			return
		}
		invoice.ExpiredDate = invoice.ExpiredDate.UTC()
		invoice.DontNotifyUntil = invoice.DontNotifyUntil.UTC()
		debtod.Invoices = append(debtod.Invoices, invoice)
	}

	return
}

func ScanDebtods(debtodRows *sql.Rows, debtod *entity.Debtod) (err error) {
	var personType string

	err = debtodRows.Scan(&debtod.Id, &debtod.Name, &debtod.Surname, &debtod.BussinessName, &debtod.Observation, &debtod.CPF_CNPJ, &personType)
	if err != nil {
		return
	}

	DeterminePersonType(debtod, &personType)

	return
}

func DeterminePersonType(debtod *entity.Debtod, personType *string) {
	switch strings.ToUpper(*personType) {
	case "PF":
		debtod.PersonType = 1
	case "PJ":
		debtod.PersonType = 2
	case "FOREIGN":
		debtod.PersonType = 3
	}
}

func ScanDebtodContacts(db *sql.DB, debtod *entity.Debtod) (err error) {
	var contact entity.Contact
	err = ScanNumbers(db, debtod, &contact)
	if err != nil {
		return
	}
	err = ScanEmails(db, debtod, &contact)
	if err != nil {
		return
	}
	debtod.Contacts = contact
	return
}

func ScanNumbers(db *sql.DB, debtod *entity.Debtod, contact *entity.Contact) (err error) {
	sqlDebtodNumbers := `
		SELECT number
		FROM tbl_debtods_contact_numbers
		WHERE fk_debtod_id = $1
		ORDER BY number DESC
	`
	numberRows, err := db.Query(sqlDebtodNumbers, &debtod.Id)
	if err != nil {
		return
	}

	for numberRows.Next() {
		var number string
		err = numberRows.Scan(&number)
		if err != nil {
			return
		}
		numberParsed, _ := strconv.Atoi(number)
		contact.Numbers = append(contact.Numbers, numberParsed)
	}

	return
}

func ScanEmails(db *sql.DB, debtod *entity.Debtod, contact *entity.Contact) (err error) {
	sqlDebtodEmails := `
		SELECT email
		FROM tbl_debtods_contact_emails
		WHERE fk_debtod_id = $1
		ORDER BY email ASC
	`
	emailRows, err := db.Query(sqlDebtodEmails, &debtod.Id)
	if err != nil {
		return
	}

	for emailRows.Next() {
		var email string
		err = emailRows.Scan(&email)
		if err != nil {
			return
		}
		contact.Emails = append(contact.Emails, email)
	}

	return
}

func ScanDebtodAddresses(db *sql.DB, debtod *entity.Debtod) (err error) {
	sqlDebtodAddresses := `
		SELECT street, number, zipcode, neighborhood, observation, description
		FROM tbl_debtods_addresses
		WHERE fk_debtod_id = $1
	`

	addressRows, err := db.Query(sqlDebtodAddresses, &debtod.Id)
	if err != nil {
		return
	}

	var addresses []entity.Address
	for addressRows.Next() {
		var a entity.Address
		err = addressRows.Scan(&a.Street, &a.Number, &a.Zipcode, &a.Neighborhood, &a.Observation, &a.Description)
		if err != nil {
			return
		}
		addresses = append(addresses, a)
	}

	debtod.Addresses = addresses

	return
}

func ScanDebtodInvoices(db *sql.DB, debtod *entity.Debtod) (err error) {
	sqlDebtodInvoices := `
		SELECT 
			id,
			invoice_identification,
			invoice_status,
			delay_notified,
			dont_notify_until,
			observation,
			description,
			expired_date          
		FROM tbl_debtods_invoices
		WHERE fk_debtod_id = $1
	`
	invoiceRows, err := db.Query(sqlDebtodInvoices, debtod.Id)
	if err != nil {
		return
	}

	var invoices []entity.Invoice
	for invoiceRows.Next() {
		var i entity.Invoice
		invoiceRows.Scan(&i.Id, &i.Identification, &i.Status, &i.DelayNotified, &i.DontNotifyUntil, &i.Observation, &i.Description, &i.ExpiredDate)
		i.ExpiredDate = i.ExpiredDate.UTC()
		i.DontNotifyUntil = i.DontNotifyUntil.UTC()
		invoices = append(invoices, i)
	}

	debtod.Invoices = invoices

	return
}
