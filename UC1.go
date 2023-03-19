package main

import "fmt"

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

func main() {
	contact := Contact{
		FirstName: "Goutham",
		LastName:  "Y",
		Address:   "RT Nagar Banglore",
		City:      "Banglore",
		State:     "Karnataka",
		Zip:       "543211",
		Phone:     "9876543211",
		Email:     "goutham@gmail.com",
	}

	fmt.Printf("Created contact details are : %+v\n", contact)
}
