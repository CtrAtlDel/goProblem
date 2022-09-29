package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
		case time.Monday:
			fmt.Println("Monday")
		
		case time.Friday:
			fmt.Println("Friday")
	}
}