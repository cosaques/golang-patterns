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
	concreteBeverage beverage
}

func (b *abstractBeverage) PrepareRecipe() {
	b.boilWater()
	b.concreteBeverage.brew()
	b.pourInCup()
	if b.concreteBeverage.customerWantsCondiments() {
		b.concreteBeverage.addCondiments()
	}
}

func (b *abstractBeverage) boilWater() {
	fmt.Println("Boiling water...")
}

func (b *abstractBeverage) pourInCup() {
	fmt.Println("Pouring in cups...")
}

// Hook method
func (b *abstractBeverage) customerWantsCondiments() bool {
	return true
}
