package entity

import (
	"errors"
	"time"
)

type Invoice struct {
	Id              string
	Identification  string
	Status          bool
	DelayNotified   bool
	DontNotifyUntil time.Time
	Observation     string
	Description     string
	ExpiredDate     time.Time
}

func NewInvoice(invoice *Invoice) (newInvoice *Invoice, err error) {
	err = invoice.Validate()
	if err != nil {
		return
	}
	newInvoice = &Invoice{
		Id:              invoice.Id,
		Identification:  invoice.Identification,
		Status:          invoice.Status,
		DelayNotified:   invoice.DelayNotified,
		DontNotifyUntil: invoice.DontNotifyUntil,
		Observation:     invoice.Observation,
		Description:     invoice.Description,
		ExpiredDate:     invoice.ExpiredDate,
	}

	return
}

func (i *Invoice) Validate() (err error) {
	if i.Id == "" {
		return errors.New("invoice id must be provided")
	}

	if i.Identification == "" {
		return errors.New("invoice identification must be provided")
	}

	return
}
