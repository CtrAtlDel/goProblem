package main

import (
	"fmt"
)

type SubBlock struct {
	str     string
	integer int
}

type Block struct {
	block   SubBlock
	str     float32
	integer float32
}

func main() {
	var obj Block
	fmt.Printf("%T\n", obj.block.integer)
	fmt.Printf("%T\n", obj.integer)
}
