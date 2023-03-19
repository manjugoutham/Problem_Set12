package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	FirstName, LastName, Address, City, State, Zip, Phone, Email string
}

type AddressBook struct {
	Contacts []Contact
}

func (a *AddressBook) FindContact(firstName, lastName string) (*Contact, int) {
	for i, c := range a.Contacts {
		if c.FirstName == firstName && c.LastName == lastName {
			return &c, i
		}
	}
	return nil, -1
}

func main() {
	addressBook := AddressBook{}

	addressBook.Contacts = append(addressBook.Contacts, Contact{
		FirstName: "Allen",
		LastName:  "Rahul",
		Address:   "banglore",
		City:      "mulbagal",
		State:     "karnataka",
		Zip:       "560040",
		Phone:     "9876543211",
		Email:     "abc@gmail.com",
	})
	addressBook.Contacts = append(addressBook.Contacts, Contact{
		FirstName: "Goutham",
		LastName:  "Y",
		Address:   "banglore",
		City:      "mulbagal",
		State:     "karnataka",
		Zip:       "560040",
		Phone:     "9876543211",
		Email:     "abc@gmail.com",
	})

	reader := bufio.NewReader(os.Stdin)

	// Prompt for contact name to edit
	fmt.Print("Enter the first name of the contact to edit: ")
	firstName, _ := reader.ReadString('\n')
	fmt.Print("Enter the last name of the contact to edit: ")
	lastName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	contact, index := addressBook.FindContact(firstName, lastName)
	if contact == nil {
		fmt.Println("Contact details are not found? Please check the names")
		return
	}

	fmt.Println("Current contact details :")
	fmt.Printf("%+v\n", *contact)

	fmt.Println("Enter new contact details:")

	fmt.Print("Enter the First Name: ")
	newFirstName, _ := reader.ReadString('\n')
	fmt.Print("Enter the Last Name: ")
	newLastName, _ := reader.ReadString('\n')
	fmt.Print("Enter the Address name : ")
	newAddress, _ := reader.ReadString('\n')
	fmt.Print("Enter the City name : ")
	newCity, _ := reader.ReadString('\n')
	fmt.Print("Enter the State name : ")
	newState, _ := reader.ReadString('\n')
	fmt.Print("Enter the Zip: ")
	newZip, _ := reader.ReadString('\n')
	fmt.Print("Enter the Phone Number: ")
	newPhone, _ := reader.ReadString('\n')
	fmt.Print("Enter the Email: ")
	newEmail, _ := reader.ReadString('\n')

	contact.FirstName = strings.TrimSpace(newFirstName)
	contact.LastName = strings.TrimSpace(newLastName)
	contact.Address = strings.TrimSpace(newAddress)
	contact.City = strings.TrimSpace(newCity)
	contact.State = strings.TrimSpace(newState)
	contact.Zip = strings.TrimSpace(newZip)
	contact.Phone = strings.TrimSpace(newPhone)
	contact.Email = strings.TrimSpace(newEmail)

	addressBook.Contacts[index] = *contact

	fmt.Printf("Contact updated:\n%+v\n", contact)
}
