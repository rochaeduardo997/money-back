package repository

import (
	"database/sql"
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

	defer dr.db.Close()

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
