package main

import (
	"fmt"

	"github.com/cosaques/patterns/algorithms/beverages"
)

func main() {
	fmt.Println("\nTEA")
	tea := beverages.NewTea()
	tea.PrepareRecipe()

	fmt.Println("\nCOFFEE")
	coffee := beverages.NewCoffee()
	coffee.PrepareRecipe()
}
