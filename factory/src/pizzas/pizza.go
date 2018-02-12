package pizzas

import (
	"fmt"

	"github.com/cosaques/patterns/factory/src/ingredients"
)

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type basePizza struct {
	name      string
	cheese    ingredients.Cheese
	clams     ingredients.Clams
	dough     ingredients.Dough
	pepperoni ingredients.Pepperoni
	sauce     ingredients.Sauce
	veggies   []ingredients.Veggie
}

func (p *basePizza) Bake() {
	fmt.Println("Pizza is baking now...")
}

func (p *basePizza) Cut() {
	fmt.Println("Pizza is cut...")
}

func (p *basePizza) Box() {
	fmt.Println("Pizza is in the box...")
}
