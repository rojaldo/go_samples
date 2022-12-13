package person

import "fmt"

type Person struct {
	name string
	age  int
}

// oveeride the default string method
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}


func (p *Person) SetName(name string) {
	if name != "" {
		p.name = name		
	}
}

func (p *Person) SetAge(age int) {
	p.age = age
}