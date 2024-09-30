package main

import "fmt"

func division(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return num1 / num2, nil
}
func main() {
	fmt.Println("stated Error Handling in the function ")
	result, err := division(60, 0)
	if err != nil {
		fmt.Println("Error Handling : ", err)
	}
	fmt.Println("Division Result", result)

}
