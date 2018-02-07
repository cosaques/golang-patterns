package condiments

import decorator "github.com/cosaques/patterns/decorator/src"

func AddMocha(b decorator.Beverage) decorator.Beverage {
	return &condiment{
		beverage:    b,
		cost:        15,
		description: "Mocha"}
}
