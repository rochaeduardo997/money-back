package test

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/rochaeduardo997/money-back/internal/Debtod/entity"
	repository "github.com/rochaeduardo997/money-back/internal/Debtod/infra"
)

func InstanceDB() (db *sql.DB) {
	db, err := sql.Open("postgres", "host=172.20.0.2 port=5432 user=psql_user password=1234512345 dbname=db_test sslmode=disable")
	if err != nil {
		panic("Failed on startup database")
	}
	CleanupDB(db)
	return
}

func CleanupDB(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE tbl_debtods CASCADE")
	if err != nil {
		panic("Failed on cleanup database" + err.Error())
	}
}

func MockAddress() (address *entity.Address) {
	address = &entity.Address{
		Street:       "street",
		Number:       944,
		Zipcode:      63333333,
		Neighborhood: "neighborhood",
		Observation:  "observation",
		Description:  "description",
	}
	return
}

func MockInvoice() (invoice *entity.Invoice) {
	invoice = &entity.Invoice{
		Id:              "uuid",
		Identification:  "identification",
		Status:          true,
		DelayNotified:   false,
		DontNotifyUntil: time.Date(2022, 02, 20, 12, 30, 0, 0, time.UTC),
		Observation:     "observation",
		Description:     "description",
		ExpiredDate:     time.Date(2022, 01, 20, 12, 30, 0, 0, time.UTC),
	}
	return
}

func MockContact() (contact *entity.Contact) {
	contact = &entity.Contact{
		Numbers: []int{54321432143, 12345123412},
		Emails:  []string{"email1", "email2"},
	}
	return
}

func MockDebtod() (debtod *entity.Debtod) {
	contacts := MockContact()
	address := MockAddress()
	invoice := MockInvoice()
	debtod = &entity.Debtod{
		Id:            "uuid",
		Name:          "name",
		Surname:       "surname",
		BussinessName: "bussiness_name",
		Observation:   "observation",
		CPF_CNPJ:      "cpf_cnpj",
		PersonType:    1,
		Contacts:      *contacts,
		Addresses:     []entity.Address{*address},
		Invoices:      []entity.Invoice{*invoice},
	}
	return
}

func Test_GivenDebtod_WhenInsertNewDebtod_ThenReceiveDebtodInstance(t *testing.T) {
	db := InstanceDB()

	debtodRepository := repository.NewDebtodRepository(db)
	defer debtodRepository.CloseDB()

	given := MockDebtod()
	got, err := debtodRepository.Save(given)
	assert.Nil(t, err)
	assert.Equal(t, given, got)
}

func Test_GivenDebtods_WhenListDebtods_ThenReceiveDebtodInstances(t *testing.T) {
	db := InstanceDB()

	debtodRepository := repository.NewDebtodRepository(db)
	defer debtodRepository.CloseDB()

	given1 := MockDebtod()
	given2 := MockDebtod()

	_, _ = debtodRepository.Save(given1)
	_, _ = debtodRepository.Save(given2)

	got, err := debtodRepository.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, given1, got[0])
	assert.Equal(t, given2, got[1])
}

func Test_GivenDebtods_WhenListDebtodById_ThenReceiveDebtodInstance(t *testing.T) {
	db := InstanceDB()

	debtodRepository := repository.NewDebtodRepository(db)
	defer debtodRepository.CloseDB()

	given1 := MockDebtod()
	given2 := MockDebtod()

	_, _ = debtodRepository.Save(given1)
	_, _ = debtodRepository.Save(given2)

	got, err := debtodRepository.GetBy(given1.Id)
	assert.Nil(t, err)
	assert.Equal(t, given1, got)
}
