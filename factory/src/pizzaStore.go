package factory

import "github.com/cosaques/patterns/factory/src/pizzas"

type pizzaStore interface {
	OrderPizza(name string)
	// Factory method
	CreatePizza(name string) pizzas.Pizza
}

type pizzaStoreBase struct {
	store pizzaStore
}

func (s *pizzaStoreBase) OrderPizza(name string) {
	pizza := s.store.CreatePizza(name)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
}

type nyPizzaStore struct {
	*pizzaStoreBase
}

func (s *nyPizzaStore) CreatePizza(name string) pizzas.Pizza {
	switch name {
	case "cheese":
		return new(pizzas.CheesePizza)
	case "pepper":
		return new(pizzas.PepperoniPizzza)
	default:
		panic("no pizza found")
	}
}

func NewNYPizzaStore() pizzaStore {
	s := new(nyPizzaStore)
	s.pizzaStoreBase = &pizzaStoreBase{store: s}

	return s
}
