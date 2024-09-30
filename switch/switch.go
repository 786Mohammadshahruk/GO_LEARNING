package main

import "fmt"

func main() {
	fmt.Println("Welcome to Switch Statement")

	//condition check
	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not ")
	}

	// condition statement check
	temp := 20

	switch {
	case temp > 10:
		fmt.Println("Hot")
	case temp < 0:
		fmt.Println("cold")
	default:
		fmt.Println("Normal ")
	}

	//Multiple condition check
	month := "january"

	switch month {
	case "january", "february", "march":
		fmt.Println("Hot")
		fmt.Println("Hot")
	case "april", "may", "june":
		fmt.Println("cold")
	default:
		fmt.Println("Normal ")
	}
}
