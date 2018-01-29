package ducks

import (
	"fmt"
)

type flyBehavior interface {
	fly()
}

type simpleFly struct{}

func (b simpleFly) fly() {
	fmt.Println("Flying.... youpi !!!!")
}

type noFly struct{}

func (b noFly) fly() {
	fmt.Println("Sorry.... I cannot fly !!!!")
}
