package factories

import "github.com/cosaques/patterns/factory/ingredients"

type NyPizzaIngredientFactory struct{}

func (f *NyPizzaIngredientFactory) CreateDough() ingredients.Dough {
	return &ingredients.ThinCrustDough{}
}

func (f *NyPizzaIngredientFactory) CreateCheese() ingredients.Cheese {
	return &ingredients.RegianoCheese{}
}

func (f *NyPizzaIngredientFactory) CreateClams() ingredients.Clams {
	return &ingredients.FreshClams{}
}

func (f *NyPizzaIngredientFactory) CreatePepperoni() ingredients.Pepperoni {
	return &ingredients.BasePepperoni{}
}

func (f *NyPizzaIngredientFactory) CreateSauce() ingredients.Sauce {
	return &ingredients.MarinaraSauce{}
}

func (f *NyPizzaIngredientFactory) CreateVeggies() []ingredients.Veggie {
	return []ingredients.Veggie{&ingredients.Onion{}, &ingredients.Garlic{}, &ingredients.MushRoom{}}
}
