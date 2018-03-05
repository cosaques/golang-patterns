package main

import "github.com/cosaques/patterns/algorithms/beverages"

func main() {
	tea := beverages.NewTea()
	tea.PrepareRecipe()

	coffee := beverages.NewCoffee()
	coffee.PrepareRecipe()
}
