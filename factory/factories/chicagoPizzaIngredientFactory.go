package factories

import "github.com/cosaques/patterns/factory/ingredients"

type ChicagoPizzaIngredientFactory struct{}

func (f *ChicagoPizzaIngredientFactory) CreateDough() ingredients.Dough {
	return &ingredients.ThickCrustDough{}
}

func (f *ChicagoPizzaIngredientFactory) CreateCheese() ingredients.Cheese {
	return &ingredients.MozzarellaCheese{}
}

func (f *ChicagoPizzaIngredientFactory) CreateClams() ingredients.Clams {
	return &ingredients.FrozenClams{}
}

func (f *ChicagoPizzaIngredientFactory) CreatePepperoni() ingredients.Pepperoni {
	return &ingredients.BasePepperoni{}
}

func (f *ChicagoPizzaIngredientFactory) CreateSauce() ingredients.Sauce {
	return &ingredients.PlumTomatoSauce{}
}

func (f *ChicagoPizzaIngredientFactory) CreateVeggies() []ingredients.Veggie {
	return []ingredients.Veggie{&ingredients.Garlic{}, &ingredients.RedPepper{}}
}
