package condiments

import "github.com/cosaques/patterns/decorator/beverages"

func AddMocha(b beverages.Beverage) beverages.Beverage {
	return &condiment{
		beverage:    b,
		cost:        15,
		description: "Mocha"}
}
