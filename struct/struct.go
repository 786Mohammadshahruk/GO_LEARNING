package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	MobileNumber string
	Email        string
}
type Address struct {
	HouseNo int
	State   string
	Pin     int64
}

type Employee struct {
	Person_Details   Person
	Contanct_Details Contact
	Address_Details  Address
}

func main() {
	fmt.Println("Welcome to Struct")

	var raj Person
	fmt.Println(raj)
	// Initializing data to Person

	raj.FirstName = "Raj"
	raj.LastName = "Kumar"
	raj.Age = 20
	fmt.Println(raj)

	// 2nd method to create and Initialize
	prem := Person{
		FirstName: "Pream",
		LastName:  "Kumar",
		Age:       22,
	}
	fmt.Println(prem)

	// 3nd method to create and Initialize
	person3 := new(Person)
	person3.FirstName = "Shahruk"
	person3.LastName = " Khan"
	person3.Age = 55
	fmt.Println(person3)
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Println("======================================================")

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Mohammad",
		LastName:  "Sahil",
		Age:       18,
	}
	employee1.Contanct_Details = Contact{
		MobileNumber: "8765432109",
		Email:        "mshahruk57@gmail.com",
	}
	employee1.Address_Details = Address{
		HouseNo: 102,
		State:   "Jharkhand",
		Pin:     833102,
	}
	fmt.Println("======================================================", employee1)
}
