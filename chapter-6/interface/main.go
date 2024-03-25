package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type Home struct {
	Address string
	Year    int
}

func (h Home) String() string {
	return fmt.Sprintf("Address: %s, Year: %d", h.Address, h.Year)
}

type Pet struct {
	Name    string
	Species string
}

func (p Pet) String() string {
	return fmt.Sprintf("Name: %s, Species: %s", p.Name, p.Species)
}

func A(s Stringer) {
	switch a := s.(type) {
	case Person:
		fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
	case Home:
		fmt.Sprintf("Address: %s, Year: %d", h.Address, h.Year)
	case Pet:
		fmt.Sprintf("Name: %s, Species: %s", p.Name, p.Species)
	}
}

func main() {

}
