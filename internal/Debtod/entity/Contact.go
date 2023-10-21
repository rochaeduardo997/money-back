package entity

import "errors"

type Contact struct {
	Numbers []int
	Emails  []string
}

func NewContact(contact *Contact) (newContact *Contact, err error) {
	err = contact.Validate()
	if err != nil {
		return
	}
	newContact = &Contact{
		Numbers: contact.Numbers,
		Emails:  contact.Emails,
	}
	return
}

func (c *Contact) Validate() (err error) {
	if len(c.Numbers) == 0 {
		return errors.New("at least one number must be provided")
	}

	return err
}
