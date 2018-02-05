package main

import (
	"fmt"
)

type Duck interface {
	Display()
	PerformFly()
}

type abstractDuck struct {
	Duck
	flyBehavior
}

func (d abstractDuck) Display() {
	fmt.Println("I'm a duck !")
}
func (d abstractDuck) PerformFly() {
	d.flyBehavior()
}

type MallardDuck struct{ abstractDuck }

func NewMallardDuck() *MallardDuck {
	d := MallardDuck{abstractDuck{}}
	d.abstractDuck.Duck = d
	d.abstractDuck.flyBehavior = simpleFly
	return &d
}

type RubberDuck struct{ abstractDuck }

func NewRubberDuck() *RubberDuck {
	d := RubberDuck{abstractDuck{}}
	d.abstractDuck.Duck = d
	d.abstractDuck.flyBehavior = noFly
	return &d
}
