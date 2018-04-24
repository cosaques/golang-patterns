package pizzaStores

import "github.com/cosaques/patterns/factory/pizzas"

type pizzaStore interface {
	OrderPizza(name string)
	createPizza(name string) pizzas.Pizza
}

type pizzaStoreBase struct {
	store pizzaStore
}

func (s *pizzaStoreBase) OrderPizza(name string) {
	pizza := s.store.createPizza(name)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
}
