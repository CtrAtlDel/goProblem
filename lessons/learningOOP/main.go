package main

import "fmt"

type Parent struct{}

func (p *Parent) Print() {
	fmt.Println("Parent")
}

type Child struct{
	parent Parent
}

func (c *Child) Print(){
	fmt.Println("Child")
}

func main(){
	var x Child
	x.Print() //print Child
}