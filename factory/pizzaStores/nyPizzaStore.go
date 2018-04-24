package pizzaStores

import (
	"github.com/cosaques/patterns/factory/factories"
	"github.com/cosaques/patterns/factory/pizzas"
)

type nyPizzaStore struct {
	*pizzaStoreBase
}

func (s *nyPizzaStore) createPizza(name string) pizzas.Pizza {
	var pizza pizzas.Pizza
	ingredientFactory := new(factories.NyPizzaIngredientFactory)

	switch name {
	case "cheese":
		pizza = pizzas.NewCheesePizza("New York style cheese pizza", ingredientFactory)
	case "clam":
		pizza = pizzas.NewClamPizza("New York style clam pizza", ingredientFactory)
	default:
		panic("no pizza found")
	}

	return pizza
}

func NewNyPizzaStore() pizzaStore {
	var s = new(nyPizzaStore)
	s.pizzaStoreBase = &pizzaStoreBase{store: s}
	return s
}
