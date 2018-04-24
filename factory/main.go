package main

import "github.com/cosaques/patterns/factory/pizzaStores"

func main() {
	store1 := pizzaStores.NewNyPizzaStore()
	store1.OrderPizza("cheese")

	store2 := pizzaStores.NewChicagoPizzaStore()
	store2.OrderPizza("clam")
}
