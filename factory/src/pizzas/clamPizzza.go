package pizzas

import (
	"fmt"

	"github.com/cosaques/patterns/factory/src/factories"
)

type clamPizza struct {
	ingredientFactory factories.PizzaIngredientFactory
	*basePizza
}

func (p *clamPizza) Prepare() {
	fmt.Printf("Preparing %s...\n", p.name)

	p.dough = p.ingredientFactory.CreateDough()
	fmt.Printf("-> Adding %T...\n", p.dough)

	p.sauce = p.ingredientFactory.CreateSauce()
	fmt.Printf("-> Adding %T...\n", p.sauce)

	p.cheese = p.ingredientFactory.CreateCheese()
	fmt.Printf("-> Adding %T...\n", p.cheese)

	p.clams = p.ingredientFactory.CreateClams()
	fmt.Printf("-> Adding %T...\n", p.clams)
}

func NewClamPizza(name string, ingredientFactory factories.PizzaIngredientFactory) *clamPizza {
	var pizza = new(clamPizza)
	pizza.basePizza = new(basePizza)
	pizza.ingredientFactory = ingredientFactory
	pizza.name = name

	return pizza
}
