package main

import "fmt"

func main() {
	fmt.Println("Welcom to Slice")

	name := []int{1, 2, 3, 4, 5}
	//fmt.Printf("%T", name)

	numbers := append(name, 0, 3, 5, 9, 2, 8)
	fmt.Println(numbers)

	var name1 = []string{}
	fmt.Println(name1)

	// New Initilization
	stock := make([]int, 5)
	fmt.Println(cap(stock))
	fmt.Println(len(stock))

}
