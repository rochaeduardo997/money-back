package entity

import (
	"errors"

	"github.com/google/uuid"
)

type EPersonType uint8

const (
	PF EPersonType = 1 << iota
	PJ
	FOREIGN
)

type Debtod struct {
	Id            string
	Name          string
	Surname       string
	BussinessName string
	CPF_CNPJ      string
	PersonType    EPersonType
	Contacts      Contact
}

func NewDebtod(debtod *Debtod) (newDebtod *Debtod, err error) {
	debtod.Id = uuid.NewString()
	err = debtod.Validate()
	if err != nil {
		return
	}
	newDebtod = &Debtod{
		Id:            debtod.Id,
		Name:          debtod.Name,
		Surname:       debtod.Surname,
		BussinessName: debtod.BussinessName,
		CPF_CNPJ:      debtod.CPF_CNPJ,
		PersonType:    debtod.PersonType,
		Contacts:      debtod.Contacts,
	}

	return newDebtod, err
}

func (d *Debtod) Validate() (err error) {
	if d.Id == "" {
		return errors.New("id must be provided")
	}
	if d.Name == "" {
		return errors.New("name must be provided")
	}
	if d.Surname == "" {
		return errors.New("surname must be provided")
	}
	if d.CPF_CNPJ == "" {
		return errors.New("cpf_cnpj must be provided")
	}
	if d.PersonType == 0 {
		return errors.New("person type must be provided")
	}

	err = d.Contacts.Validate()
	if err != nil {
		return
	}

	return nil
}