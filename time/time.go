package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println("Current Time", current_time)
	formatted_time := current_time.Format("02-01-2006, Monday, 15:04:05")
	fmt.Println("Formatted Time", formatted_time)

	// Time conversion

	str_time := "30-09-2024"
	converted_date_time, _ := time.Parse("02-01-2006", str_time)
	fmt.Println("Converted Date", converted_date_time)
	//Add 1 more day in current date
	newDate := current_time.Add(24 * time.Hour)
	fmt.Println(newDate)
}
