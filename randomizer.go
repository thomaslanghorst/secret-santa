package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Randomize(contacts []*Contact) map[*Contact]*Contact {

	pickedContacts := make(map[*Contact]*Contact)

	for _, pickingContact := range contacts {

		contactOK := false
		pickedContact := &Contact{}

		for !contactOK {
			pickedContact = pickContact(contacts, pickingContact, pickedContacts)
			contactOK = validatePick(pickedContacts, pickedContact, pickingContact)
		}

		pickedContacts[pickedContact] = pickingContact
	}

	return pickedContacts
}

func pickContact(contacts []*Contact, pickingContact *Contact, pickedContacts map[*Contact]*Contact) *Contact {
	nr := randomNumber(0, len(contacts)-1)
	pickedContact := contacts[nr]
	return pickedContact
}

func validatePick(pickedContacts map[*Contact]*Contact, pickedContactName, pickingContactName *Contact) bool {

	// contact picked himself
	if pickedContactName.Name == pickingContactName.Name {
		return false
	}

	// contact already picked
	if pickedContacts[pickedContactName] != nil {
		return false
	}

	return true
}

func randomNumber(min, max int) uint32 {
	return uint32(rand.Int31n(int32(max+1)-int32(min)) + int32(min))
}
