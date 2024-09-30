package main

import "fmt"

func myFunction() {
	fmt.Println("Sample Function")
}

func main() {
	fmt.Println("Main Function")
	myFunction()

	ans := add1(40, 42)
	fmt.Println("Answer :", ans)

	result := add2(10, 20)
	fmt.Println("Result :", result)
}

func add1(a int, b int) int {
	return a + b
}
func add2(a int, b int) (result int) {
	result = a + b
	return
}
