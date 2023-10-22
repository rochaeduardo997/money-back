package entity

import "errors"

type Address struct {
	Street       string
	Number       uint16
	Zipcode      uint32
	Neighborhood string
	Observation  string
	Description  string
}

func NewAddress(address *Address) (newAddress *Address, err error) {
	err = address.Validate()
	if err != nil {
		return
	}
	newAddress = &Address{
		Street:       address.Street,
		Number:       address.Number,
		Zipcode:      address.Zipcode,
		Neighborhood: address.Neighborhood,
		Observation:  address.Observation,
		Description:  address.Description,
	}
	return
}

func (a *Address) Validate() (err error) {
	if a.Street == "" {
		return errors.New("street must be provided")
	}
	if a.Number == 0 {
		return errors.New("number must be provided")
	}
	if a.Zipcode == 0 {
		return errors.New("zipcode must be provided")
	}

	return
}
