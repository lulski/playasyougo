package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Grey",
		NumberOfTeeth: 24,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal name is", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Dog) Says() string {
	return d.Name
}

func (d *Gorilla) NumberOfLegs() int {
	return d.NumberOfTeeth
}

func (d *Gorilla) Says() string {
	return d.Name
}
