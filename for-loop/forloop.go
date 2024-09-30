package main

import "fmt"

func main() {
	fmt.Println("For Loops")

	for i := 0; i < 10; i++ {
		fmt.Println("Index : ", i)
	}

	//Infinite Loop

	counter := 1
	for {
		fmt.Println("Infinite Loops")
		counter++
		if counter == 3 {
			break
		}
	}

	// Range prove index and value
	numbers := []int{11, 22, 34, 41, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, and  Value Is %d\n", index, value)
	}

	data := "Hello Shahruk"
	for index, char := range data {
		fmt.Printf("Index: %d, and  Char Is %c\n", index, char)
	}

}
