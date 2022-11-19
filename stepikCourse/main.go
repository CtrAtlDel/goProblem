package main

import "fmt"

func main() {
	var n, b int
	sum := 0
	
	for i := 0; i < n; i++ {
		fmt.Scan(&b)

		if (b > 9) && (b < 100) && (b%8 == 0) {
			sum += sum + b
		}
	}
	fmt.Println(sum)
}
