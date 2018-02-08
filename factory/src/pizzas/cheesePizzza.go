package pizzas

import (
	"fmt"
)

type CheesePizza struct{}

func (p *CheesePizza) Prepare() {
	fmt.Println("Mmmm... preparing nice cheese pizza")
}

func (p *CheesePizza) Bake() {
	fmt.Println("Cheese pizza is baking now")
}

func (p *CheesePizza) Cut() {
	fmt.Println("Cheese pizza is cut... mm.. how much cheese on the knife")
}

func (p *CheesePizza) Box() {
	fmt.Println("Cheese pizza is in the box")
}
