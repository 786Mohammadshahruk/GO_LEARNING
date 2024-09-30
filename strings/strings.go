package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string = "apple,orange,banana"
	var strArray []string = strings.Split(data, ",")
	fmt.Println(strArray)
}
