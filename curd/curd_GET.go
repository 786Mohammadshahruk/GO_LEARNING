package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	USER_ID   int    `json:"userId"`
	Id        int    `json:'"id"`
	TITLE     string `json:"title"`
	COMPLETED bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning Curd")

	response, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if error != nil {
		fmt.Println("Getting Error While Calling ToDos")
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error In Getting Response")
		return
	}
	// response_byte, error := ioutil.ReadAll(response.Body)
	// fmt.Println("Response Byte", string(response_byte))

	var user User
	error1 := json.NewDecoder(response.Body).Decode(&user)
	if error1 != nil {
		fmt.Println("Getting Error While Decoding ToDos")
		return
	}
	fmt.Println("USER : ", user)
}
