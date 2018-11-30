package strategy

import (
	"fmt"
)

type Duck interface {
	Swim()
	Display()
	PerformFly()
}

type baseDuck struct {
	flyBehavior
}

func (d baseDuck) Swim() {
	fmt.Println("Swimming...")
}
func (d baseDuck) PerformFly() {
	d.flyBehavior()
}

type MallardDuck struct{ baseDuck }

func (d MallardDuck) Display() {
	fmt.Println("I'm a MallardDuck !")
}
func NewMallardDuck() *MallardDuck {
	d := MallardDuck{baseDuck{}}
	d.baseDuck.flyBehavior = simpleFly
	return &d
}

type RubberDuck struct{ baseDuck }

func (d RubberDuck) Display() {
	fmt.Println("I'm a RubberDuck !")
}
func NewRubberDuck() *RubberDuck {
	d := RubberDuck{baseDuck{}}
	d.baseDuck.flyBehavior = noFly
	return &d
}
