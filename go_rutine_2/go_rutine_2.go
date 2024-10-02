package main

import (
	"fmt"
	"sync"
)

func worker(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started,\n", value)

	fmt.Printf("Worker %d End,\n", value)

}
func main() {
	//fmt.Println("Welcome to go rutine")

	var wg sync.WaitGroup
	//starting 3 worker gorutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // incrementing the wait group counter
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Worker Task completed")
}
