package test

import (
	"testing"
	"time"

	"github.com/rochaeduardo997/money-back/internal/Debtod/entity"
	"github.com/stretchr/testify/assert"
)

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

func Test_GivenEmptyId_WhenCreateNewInvoice_ThenReceiveError(t *testing.T) {
	got := entity.Invoice{}
	assert.EqualError(t, got.Validate(), "invoice id must be provided")
}

func Test_GivenEmptyIdentification_WhenCreateNewInvoice_ThenReceiveError(t *testing.T) {
	got := entity.Invoice{Id: "uuid"}
	assert.EqualError(t, got.Validate(), "invoice identification must be provided")
}

func Test_GivenInvoice_WhenCreateNewInvoice_ThenReceiveNormalInvoiceInstance(t *testing.T) {
	invoice := &entity.Invoice{
		Id:              "uuid",
		Identification:  "identification",
		Status:          true,
		DelayNotified:   false,
		DontNotifyUntil: time.Date(2022, 02, 20, 12, 30, 0, 0, time.UTC),
		Observation:     "observation",
		Description:     "description",
		ExpiredDate:     time.Date(2022, 01, 20, 12, 30, 0, 0, time.UTC),
	}
	got, err := entity.NewInvoice(invoice)
	assert.Nil(t, err)
	assert.EqualValues(t, got, invoice)
}
