package main

import "fmt"

type triangle struct{
	Name string
	Weight, Height float32
}

func main() {
	var tr1 triangle
	var tr2 = triangle{"I'm tr.", 10, 20}
	var tr3 = triangle{Height: 12, Weight: 14, Name: "I'm trrr"}

	fmt.Println(tr1)
	fmt.Println(tr2)
	fmt.Println(tr3)
}