package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string = "apple,orange,banana"
	var strArray []string = strings.Split(data, ",")
	fmt.Println(strArray)

	str := "    Hello Let's Go          "
	fmt.Println(str)
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "Mohammad"
	str2 := "Shahruk"
	result := strings.Join([]string{str1, str2}, ",")
	fmt.Println(result)

}
