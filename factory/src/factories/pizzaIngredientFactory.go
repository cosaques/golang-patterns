package factories

import "github.com/cosaques/patterns/factory/src/ingredients"

type PizzaIngredientFactory interface {
	CreateDough() ingredients.Dough
	CreateCheese() ingredients.Cheese
	CreateClams() ingredients.Clams
	CreatePepperoni() ingredients.Pepperoni
	CreateSauce() ingredients.Sauce
	CreateVeggies() []ingredients.Veggie
}
