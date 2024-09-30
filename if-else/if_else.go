package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else Statement")

	x := 20
	if x > 6 {
		fmt.Println("Grater")
	} else if x < 5 && (x > 9) {
		fmt.Println("Smaller")
	} else {
		fmt.Println("Zero")
	}
}
