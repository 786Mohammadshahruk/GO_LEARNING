package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Say Hello")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Say Hello Function Ended")
}

func sayHi() {
	fmt.Println("Say Hiii")
}

func main() {
	fmt.Println("Hello Welcome to GO Rutine")

	go sayHello()
	sayHi()
	time.Sleep(3000 * time.Millisecond)
}
