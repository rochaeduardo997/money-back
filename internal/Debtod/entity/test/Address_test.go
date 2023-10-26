package test

import (
	"testing"

	"github.com/rochaeduardo997/money-back/internal/Debtod/entity"
	"github.com/stretchr/testify/assert"
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

func Test_GivenEmptyStreet_WhenCreateNewAddress_ThenReceiveError(t *testing.T) {
	got := entity.Address{}
	assert.EqualError(t, got.Validate(), "street must be provided")
}

func Test_GivenEmptyNumber_WhenCreateNewAddress_ThenReceiveError(t *testing.T) {
	got := entity.Address{Street: "street"}
	assert.EqualError(t, got.Validate(), "number must be provided")
}

func Test_GivenEmptyZipcode_WhenCreateNewAddress_ThenReceiveError(t *testing.T) {
	got := entity.Address{Street: "street", Number: 944}
	assert.EqualError(t, got.Validate(), "zipcode must be provided")
}

func Test_GivenAddress_WhenCreateNewAddress_ThenReceiveNormalAddressInstance(t *testing.T) {
	address := &entity.Address{
		Street:       "street",
		Number:       944,
		Zipcode:      63333333,
		Neighborhood: "neighborhood",
		Observation:  "observation",
		Description:  "description",
	}
	got, err := entity.NewAddress(address)
	assert.Nil(t, err)
	assert.EqualValues(t, got, address)
}
