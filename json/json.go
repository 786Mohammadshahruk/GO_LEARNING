package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsAudult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning JSON")
	person := Person{
		ID:       101,
		Name:     "Shahruk",
		Age:      20,
		IsAudult: true,
	}
	fmt.Println(person)

	//Convert to JSON Encoding Marshalling
	jsonDate, error := json.Marshal(person)
	if error != nil {
		fmt.Println("Error While Marshalling ")
		return
	}
	fmt.Println(string(jsonDate))

	//Unmarshalling

	var personData Person
	error1 := json.Unmarshal(jsonDate, &personData)
	if error1 != nil {
		fmt.Println("Error While Un-marshalling ")
		return
	}
	fmt.Println("Person Data : ", personData)
}
