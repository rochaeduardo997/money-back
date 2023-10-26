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
	db, err := sql.Open("postgres", "host=db port=5432 user=psql_user password=1234512345 dbname=db_test sslmode=disable")
	assert.Nil(t, err)

	debtodRepository := repository.NewDebtodRepository(db)

	given := MockDebtod()
	_, err = debtodRepository.Save(given)
	assert.Nil(t, err)
}