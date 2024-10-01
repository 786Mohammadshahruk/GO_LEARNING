package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning Web Service")
	//https://jsonplaceholder.typicode.com/todos
	response, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if error != nil {
		fmt.Println("Error While featching Data")
	}
	defer response.Body.Close()
	fmt.Printf("Response Type %T\n:", response)
	data, erro1r := ioutil.ReadAll(response.Body)
	if erro1r != nil {
		fmt.Println("Error While featching Data")
	}
	fmt.Println("Response Data:", string(data))
}
