package main;

import "fmt"

// study chanels

func sum(s []int, c chan int ) {
	sum := 0
	for _, it := range s {
		sum += it
	}
	c <- sum //chanel send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <- c, <- c // receive from c
	fmt.Println(x, y, x+y)
}