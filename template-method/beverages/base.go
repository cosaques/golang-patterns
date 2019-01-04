package beverages

import (
	"fmt"
)

type Beverage interface {
	PrepareRecipe()
}

type subBeverage interface {
	brew()
	addCondiments()
	customerWantsCondiments() bool
}

type baseBeverage struct {
	sub subBeverage
}

func (b *baseBeverage) PrepareRecipe() {
	b.boilWater()
	b.sub.brew()
	b.pourInCup()
	if b.sub.customerWantsCondiments() {
		b.sub.addCondiments()
	}
}

func (b *baseBeverage) boilWater() {
	fmt.Println("Boiling water...")
}

func (b *baseBeverage) pourInCup() {
	fmt.Println("Pouring in cups...")
}

// Hook method
func (b *baseBeverage) customerWantsCondiments() bool {
	return true
}
