package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Map")
	printMap()

}

func printMap() {

	student := make(map[string]int)
	student["Mohammad"] = 210
	student["Shahruk"] = 200
	student["Rakesh"] = 100
	student["Gaurav"] = 150
	fmt.Println(student)
	fmt.Println(student["Mohammad"])
	delete(student, "Mohammad")
	fmt.Println(student["Mohammad"])

	value, exist := student["Rakesh"]
	fmt.Println("Value : ", value)
	fmt.Println("Exist : ", exist)

	for key, value := range student {
		fmt.Printf("Map Of key: %s, and Value : %d \n", key, value)
	}

	// Diffrent way to stor values

	paerson := map[int]string{
		12: "sameer",
		50: "khan",
		1:  "ajay",
		32: "sahab",
	}

	for personId, personName := range paerson {
		fmt.Printf("Person Id Is : %d, and Person Name Is : %s \n", personId, personName)
	}

}
