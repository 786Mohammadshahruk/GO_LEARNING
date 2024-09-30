package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter user Name")
	// var name string

	// fmt.Scan(&name)
	// fmt.Println("My Name Is ", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("My Name Is ", name)
}
