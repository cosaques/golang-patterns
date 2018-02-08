package factory

import (
	"github.com/cosaques/patterns/factory/src/pizzas"
)

type simplePizzaFactory struct{}

func (f *simplePizzaFactory) createPizza(name string) pizzas.Pizza {
	switch name {
	case "cheese":
		return new(pizzas.CheesePizza)
	case "pepper":
		return new(pizzas.PepperoniPizzza)
	default:
		panic("no pizza found")
	}
}
