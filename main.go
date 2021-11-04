package main

import (
	"fmt"
	"log"
)

var (
	csvFile = "contacts.csv"
)

func main() {
	wac := NewWhatsAppClient(3, 2123, 7)
	err := wac.Login()
	if err != nil {
		log.Fatal("unable to log into whatsapp")
	}

	contacts, err := ReadContacts(csvFile)
	if err != nil {
		log.Fatal("error reading csv file")
	}

	picks := Randomize(contacts)

	for picked, picking := range picks {
		msgId, err := wac.SendMessage(picking.Number, fmt.Sprintf("You picked: %s", picked.Name))
		if err != nil {
			log.Fatal("error sending message")
		}

		fmt.Printf("Send secret santa to: %s\n", picking.Name)

		if picking.Name != "Thomas" {
			err = wac.DeleteMessage(picking.Number, msgId)
			if err != nil {
				log.Fatal("error deleting message")
			}

			fmt.Printf("Removed message for: %s\n", picking.Name)
		}
	}

	wac.Logout()
}
