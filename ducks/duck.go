package ducks

import (
	"fmt"
)

type Duck interface {
	Name() string
	Display()
	Quack()
}

type abstractDuck struct{ Duck }

func (d abstractDuck) Display() {
	fmt.Printf("Hello my name is %s \n", d.Name())
}
func (d abstractDuck) Quack() {
	fmt.Println("Kwaaaaa")
}

type MallardDuck struct{ abstractDuck }

func (d MallardDuck) Name() string {
	return "MallardDuck"
}
func NewMallardDuck() *MallardDuck {
	d := MallardDuck{abstractDuck{}}
	d.abstractDuck.Duck = d
	return &d
}

type RubberDuck struct{ abstractDuck }

func (d RubberDuck) Quack() {
	fmt.Println("Tsssss....")
}
func (d RubberDuck) Name() string {
	return "RubberDuck"
}
func NewRubberDuck() *RubberDuck {
	d := RubberDuck{abstractDuck{}}
	d.abstractDuck.Duck = d
	return &d
}
