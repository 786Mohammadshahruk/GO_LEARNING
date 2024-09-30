package main

import "fmt"

func main() {
	fmt.Println("We are Learning Array")

	var array1 = [10]int{10, 20, 30}
	fmt.Println(array1)

	array2 := [10]int{100, 200, 300}
	fmt.Println(array2)

	var array3 = [...]int{10, 20, 30, 40}
	fmt.Println(array3)

	array4 := [...]int{100, 200, 300, 400}
	fmt.Println(array4)

	//Not Initialized
	array5 := [10]int{}
	fmt.Println(array5)

	//Not Initialized
	array6 := [...]int{}
	fmt.Println(array6)

	var name [5]string
	//name[0] = "abd"
	fmt.Println(name)
	fmt.Println(len(name))

	var bo [5]bool
	fmt.Print(bo)

}
