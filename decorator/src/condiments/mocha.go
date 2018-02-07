package condiments

import decorator "github.com/cosaques/patterns/decorator/src"

type mocha struct {
	condiment
}

func AddMocha(b decorator.Beverage) *mocha {
	return &mocha{condiment{
		beverage:    b,
		cost:        15,
		description: "Mocha"}}
}
