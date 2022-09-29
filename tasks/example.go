package main

import ("fmt")

func main() {
	x := 3
	y := &x
	fmt.Print(*y)
	*y = 4
	fmt.Println(x)
}