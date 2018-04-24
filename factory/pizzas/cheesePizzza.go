package pizzas

import (
	"fmt"

	"github.com/cosaques/patterns/factory/factories"
)

type cheesePizza struct {
	ingredientFactory factories.PizzaIngredientFactory
	*basePizza
}

func (p *cheesePizza) Prepare() {
	fmt.Printf("Preparing %s...\n", p.name)

	p.dough = p.ingredientFactory.CreateDough()
	fmt.Printf("-> Adding %T...\n", p.dough)

	p.sauce = p.ingredientFactory.CreateSauce()
	fmt.Printf("-> Adding %T...\n", p.sauce)

	p.cheese = p.ingredientFactory.CreateCheese()
	fmt.Printf("-> Adding %T...\n", p.cheese)
}

func NewCheesePizza(name string, ingredientFactory factories.PizzaIngredientFactory) *cheesePizza {
	var pizza = new(cheesePizza)
	pizza.basePizza = new(basePizza)
	pizza.ingredientFactory = ingredientFactory
	pizza.name = name

	return pizza
}
