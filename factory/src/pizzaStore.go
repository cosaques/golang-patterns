package factory

type pizzaStore struct {
	pizzaFactory *simplePizzaFactory
}

func (s *pizzaStore) OrderPizza(name string) {
	pizza := s.pizzaFactory.createPizza(name)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
}

func NewPizzaStore() *pizzaStore {
	return &pizzaStore{pizzaFactory: &simplePizzaFactory{}}
}
