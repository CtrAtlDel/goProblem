package main 

import "fmt"

type animal interface {
	makeSound()
}

type cat struct{}
type dog struct{}

func (c *cat) makeSound() {
	fmt.Println("meow")
}

func (c *dog) makeSound() {
	fmt.Println("wof")
}

func main() {
	var c, d animal = &cat{}, &dog{}
	c.makeSound()
	d.makeSound()
}
