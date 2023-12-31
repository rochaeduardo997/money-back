package test

import (
	"testing"

	"github.com/rochaeduardo997/money-back/internal/Debtod/entity"
	"github.com/stretchr/testify/assert"
)

func Test_GivenDebtod_WhenCreateNewDebtod_ThenReceiveNormalDebtodInstance(t *testing.T) {
	contacts := MockContact()
	address := MockAddress()
	invoice := MockInvoice()
	debtod := &entity.Debtod{
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
		IsActive:      true,
	}
	got, err := entity.NewDebtod(debtod)
	assert.Nil(t, err)
	assert.EqualValues(t, got, debtod)
}

func Test_GivenDebtodWithoutAddress_WhenCreateNewDebtod_ThenReceiveNormalDebtodInstance(t *testing.T) {
	contacts := MockContact()
	invoice := MockInvoice()
	debtod := &entity.Debtod{
		Id:            "uuid",
		Name:          "name",
		Surname:       "surname",
		BussinessName: "bussiness_name",
		CPF_CNPJ:      "cpf_cnpj",
		PersonType:    1,
		Contacts:      *contacts,
		Invoices:      []entity.Invoice{*invoice},
		IsActive:      true,
	}
	got, err := entity.NewDebtod(debtod)
	assert.Nil(t, err)
	assert.EqualValues(t, got, debtod)
}

func Test_GivenEmptyID_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	got := entity.Debtod{}
	assert.EqualError(t, got.Validate(), "id must be provided")
}

func Test_GivenEmptyName_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	got := entity.Debtod{Id: "uuid"}
	assert.EqualError(t, got.Validate(), "name must be provided")
}

func Test_GivenEmptySurname_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	got := entity.Debtod{Id: "uuid", Name: "name"}
	assert.EqualError(t, got.Validate(), "surname must be provided")
}

func Test_GivenEmptyCPFCNPJ_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	got := entity.Debtod{
		Id:      "uuid",
		Name:    "name",
		Surname: "surname",
	}
	assert.EqualError(t, got.Validate(), "cpf_cnpj must be provided")
}

func Test_GivenEmptyPersonType_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	got := entity.Debtod{
		Id:       "uuid",
		Name:     "name",
		Surname:  "surname",
		CPF_CNPJ: "cpf_cnpj",
	}
	assert.EqualError(t, got.Validate(), "person type must be provided")
}

func Test_GivenEmptyContacts_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	got := entity.Debtod{
		Id:         "uuid",
		Name:       "name",
		Surname:    "surname",
		CPF_CNPJ:   "cpf_cnpj",
		PersonType: 1,
	}
	assert.EqualError(t, got.Validate(), "at least one number must be provided")
}

func Test_GivenEmptyAddressStreet_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	contacts := MockContact()
	address := MockAddress()
	address.Street = ""
	got := entity.Debtod{
		Id:         "uuid",
		Name:       "name",
		Surname:    "surname",
		CPF_CNPJ:   "cpf_cnpj",
		PersonType: 1,
		Contacts:   *contacts,
		Addresses:  []entity.Address{*address},
	}
	assert.EqualError(t, got.Validate(), "street must be provided")
}

func Test_GivenEmptyInvoices_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	contacts := MockContact()
	address := MockAddress()
	got := entity.Debtod{
		Id:         "uuid",
		Name:       "name",
		Surname:    "surname",
		CPF_CNPJ:   "cpf_cnpj",
		PersonType: 1,
		Contacts:   *contacts,
		Addresses:  []entity.Address{*address},
	}
	assert.EqualError(t, got.Validate(), "at least one invoice must be provided")
}

func Test_GivenEmptyInvoicesId_WhenCreateNewDebtod_ThenReceiveError(t *testing.T) {
	contacts := MockContact()
	address := MockAddress()
	invoice := MockInvoice()
	invoice.Id = ""
	got := entity.Debtod{
		Id:         "uuid",
		Name:       "name",
		Surname:    "surname",
		CPF_CNPJ:   "cpf_cnpj",
		PersonType: 1,
		Contacts:   *contacts,
		Addresses:  []entity.Address{*address},
		Invoices:   []entity.Invoice{*invoice},
	}
	assert.EqualError(t, got.Validate(), "invoice id must be provided")
}
