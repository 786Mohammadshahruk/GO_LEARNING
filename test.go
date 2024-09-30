package main

import (
	"fmt"
	"mylearning/myutils"
)

func main() {
	fmt.Println("Hello World")
	myutils.PrintMessage("Hello My Utils")
	fmt.Println("------------------------------------------------------------------------------------------------------------")

	//String
	var name string = "Shahruk"
	fmt.Println(name)

	//Integer
	var version = 10
	fmt.Println(version)

	//Float
	var dimention float64 = 10.30
	fmt.Println(dimention)

	// boolean
	var isApplicable bool = true
	isApplicable = false
	fmt.Println(isApplicable)

	const c = 10
	//c = 20   Error
	fmt.Println(c)
	fmt.Println("------------------------------------------------------------------------------------------------------------")

	person := 20
	fmt.Println(person)

	Public := "Mohammad Shahruk"
	fmt.Println(Public)

}
