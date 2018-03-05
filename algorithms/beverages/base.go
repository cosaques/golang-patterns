package beverages

import (
	"fmt"
)

type beverage interface {
	brew()
	addCondiments()
	hook()
}

type abstractBeverage struct {
	beverage
}

func (b *abstractBeverage) PrepareRecipe() {
	b.boilWater()
	b.brew()
	b.pourInCup()
	b.addCondiments()
	b.beverage.hook()
}

func (b *abstractBeverage) boilWater() {
	fmt.Println("Boiling water...")
}

func (b *abstractBeverage) pourInCup() {
	fmt.Println("Pouring in cups...")
}

func (b *abstractBeverage) hook() {
}
