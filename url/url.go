package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")

	parseUrl, error := url.Parse("https://jsonplaceholder.typicode.com/todos/1")

	if error != nil {
		fmt.Println("Error While Parshing URL ")
		return
	}
	fmt.Printf("Parse URL Type %T\n", parseUrl)

	fmt.Println("Parse URL HOST ", parseUrl.Host)
	fmt.Println("Parse URL Scheme ", parseUrl.Scheme)
	fmt.Println("Parse URL Path ", parseUrl.Path)

}
