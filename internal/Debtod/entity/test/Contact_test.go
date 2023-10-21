package entity_test

import (
	"testing"

	"github.com/rochaeduardo997/money-back/internal/Debtod/entity"
	"github.com/stretchr/testify/assert"
)

func NewContact() (contact *entity.Contact) {
	contact = &entity.Contact{
		Numbers: []int{54321432143, 12345123412},
		Emails:  []string{"email1", "email2"},
	}
	return
}

func Test_GivenEmptyNumber_WhenCreateNewContact_ThenReceiveError(t *testing.T) {
	got := entity.Contact{}
	assert.EqualError(t, got.Validate(), "at least one number must be provided")
}

func Test_GivenContact_WhenCreateNewContact_ThenReceiveNormalContactInstance(t *testing.T) {
	contact := &entity.Contact{
		Numbers: []int{54321432143, 12345123412},
		Emails:  []string{"email1", "email2"},
	}
	got, err := entity.NewContact(contact)
	assert.Nil(t, err)
	assert.EqualValues(t, got, contact)
}
