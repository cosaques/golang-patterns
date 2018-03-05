package beverages

import (
	"fmt"
)

type coffee struct {
	*abstractBeverage
}

func (b *coffee) brew() {
	fmt.Println("Brewing coffee...")
}

func (b *coffee) addCondiments() {
	fmt.Println("Adding milk and sugar...")
}

func (b *coffee) hook() {
	fmt.Println("HOOOK !!!! ...")
}

func NewCoffee() *coffee {
	c := &coffee{}
	c.abstractBeverage = &abstractBeverage{c}
	return c
}
