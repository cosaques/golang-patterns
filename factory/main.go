package main

import (
	factory "github.com/cosaques/patterns/factory/src"
)

func main() {
	store := factory.NewPizzaStore()
	store.OrderPizza("cheese")
	store.OrderPizza("pepper")
}
