package pizzas

import (
	"fmt"
)

type PepperoniPizzza struct{}

func (p *PepperoniPizzza) Prepare() {
	fmt.Println("Mmmm... preparing nice Pepper pizza")
}

func (p *PepperoniPizzza) Bake() {
	fmt.Println("Pepper pizza is baking now")
}

func (p *PepperoniPizzza) Cut() {
	fmt.Println("Pepper pizza is cut... mm.. how much pepper on the knife")
}

func (p *PepperoniPizzza) Box() {
	fmt.Println("Pepper pizza is in the box")
}
