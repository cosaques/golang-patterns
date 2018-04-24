package pizzaStores

import (
	"github.com/cosaques/patterns/factory/factories"
	"github.com/cosaques/patterns/factory/pizzas"
)

type chicagoPizzaStore struct {
	*pizzaStoreBase
}

func (s *chicagoPizzaStore) createPizza(name string) pizzas.Pizza {
	var pizza pizzas.Pizza
	ingredientFactory := new(factories.ChicagoPizzaIngredientFactory)

	switch name {
	case "cheese":
		pizza = pizzas.NewCheesePizza("Chicago style cheese pizza", ingredientFactory)
	case "clam":
		pizza = pizzas.NewClamPizza("Chicago style clam pizza", ingredientFactory)
	default:
		panic("no pizza found")
	}

	return pizza
}

func NewChicagoPizzaStore() pizzaStore {
	var s = new(chicagoPizzaStore)
	s.pizzaStoreBase = &pizzaStoreBase{store: s}
	return s
}
