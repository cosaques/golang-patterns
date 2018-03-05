package beverages

import (
	"fmt"
)

type beverage interface {
	brew()
	addCondiments()
	customerWantsCondiments() bool
}

type abstractBeverage struct {
	beverage
}

func (b *abstractBeverage) PrepareRecipe() {
	b.boilWater()
	b.brew()
	b.pourInCup()
	if b.beverage.customerWantsCondiments() {
		b.addCondiments()
	}
}

func (b *abstractBeverage) boilWater() {
	fmt.Println("Boiling water...")
}

func (b *abstractBeverage) pourInCup() {
	fmt.Println("Pouring in cups...")
}

func (b *abstractBeverage) customerWantsCondiments() bool {
	return true
}
