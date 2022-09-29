package main

import "fmt"

type MyValue interface {
	String() string
}

const (
	read = iota
	write
)

// the same

const (
	title = iota
	author = iota
)

func main() {
	fmt.Println(write)
	fmt.Println(title)
}