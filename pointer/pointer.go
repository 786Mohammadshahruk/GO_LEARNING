package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointer")

	var num int = 100

	var ptr *int = &num

	fmt.Println("Number :", num)
	fmt.Println("Address :", ptr)

	//default value of pointer

	var pointer *string
	if pointer == nil {
		fmt.Println("Data has not been assigned")
	}

	value := 100
	modifyTheValue(&value)
	fmt.Println("New Value has been assigned", value)

}

func modifyTheValue(address *int) {
	*address = *address * 2
}
