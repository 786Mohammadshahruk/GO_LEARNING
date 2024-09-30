package main

import "fmt"

func main() {
	age := 18
	name := "Shahruk"
	height := 5.7244
	fmt.Print("age : ", age, " name : ", name, " height: ", height)
	fmt.Print(" Hello World ")
	fmt.Println(" --------------------------------------------------- ")
	fmt.Printf(" Age: %d\n", age)
	fmt.Printf(" height: %.3f\n", height)

	fmt.Printf(" name: %s\n", name)
	fmt.Println(" --------------------------------------------------- ")

	fmt.Printf("Type of Age Is %T\n", age)
	fmt.Printf("Type of Name Is %T\n", name)
	fmt.Printf("Type of Height Is %T\n", height)
}
