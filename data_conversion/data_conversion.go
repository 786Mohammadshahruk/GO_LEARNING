package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 32
	fmt.Println("Number Is ", num)
	fmt.Printf("Number Type Is %T\n", num)

	var convered_data float64 = float64(num)
	fmt.Println("Converted Number Is ", convered_data)
	fmt.Printf("Converted Number Type Is %T\n", convered_data)

	//number to String
	num = 123
	str := strconv.Itoa(num)

	fmt.Println("str Number Is ", str)
	fmt.Printf("str Number Type Is %T\n", str)

	number_string := "12345"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("number_int Number Is ", number_int)
	fmt.Printf("number_int Number Type Is %T\n", number_int)

	num_string := "12345.12"
	num_int, _ := strconv.ParseFloat(num_string, 64)

	fmt.Println("num_int Number Is ", num_int)
	fmt.Printf("num_int Number Type Is %T\n", num_int)

}
