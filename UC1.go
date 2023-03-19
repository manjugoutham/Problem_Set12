package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contact struct {
	FirstName string
	LastName  string
	Address   string
	City      string
	State     string
	Zip       string
	Phone     string
	Email     string
}

type AddressBook struct {
	Contacts []Contact
}

func (a *AddressBook) AddContact(c Contact) {
	a.Contacts = append(a.Contacts, c)
}

func main() {
	addressBook := AddressBook{}

	fmt.Println("Add a new contact:")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the FirstName: ")
	firstName, _ := reader.ReadString('\n')
	fmt.Print("Enter the LastName: ")
	lastName, _ := reader.ReadString('\n')
	fmt.Print("Enter the Address: ")
	address, _ := reader.ReadString('\n')
	fmt.Print("Enter the City: ")
	city, _ := reader.ReadString('\n')
	fmt.Print("Enter the State: ")
	state, _ := reader.ReadString('\n')
	fmt.Print("Enter the Zip: ")
	zip, _ := reader.ReadString('\n')
	fmt.Print("Enter the Phone Number: ")
	phone, _ := reader.ReadString('\n')
	fmt.Print("Enter the Email: ")
	email, _ := reader.ReadString('\n')

	contact := Contact{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		City:      city,
		State:     state,
		Zip:       zip,
		Phone:     phone,
		Email:     email,
	}

	addressBook.AddContact(contact)

	fmt.Printf("Contact added:\n%+v\n", contact)
}
