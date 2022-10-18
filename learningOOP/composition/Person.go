package composition

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Inroduce() {
	fmt.Printf("Hi, I am %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Inroduce()// <-- одно и тоже
	goku.Person.Inroduce()// <-- одно и тоже
}
