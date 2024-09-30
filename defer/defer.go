package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting of the Program")
	defer fmt.Println("Middle of the Program")
	fmt.Println("end of the Program")
}
