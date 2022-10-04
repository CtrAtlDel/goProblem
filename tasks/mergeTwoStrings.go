package main

import (
	"fmt"
)

func merge(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	var s1 string
	var s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Println(merge(s1, s2))
}
