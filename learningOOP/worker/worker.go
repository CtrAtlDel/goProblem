package worker

import "fmt"

type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

func NewSaiyan(name string, power int) (*Saiyan) {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}

func (s *Saiyan) Super() {
	s.Power +=10_000
}

func main() {
	goku_1:=new (Saiyan)
	goku_2:=&Saiyan{}
	fmt.Println(goku_1.Name)
	fmt.Println(goku_2.Name)
}