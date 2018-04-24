package condiments

import "github.com/cosaques/patterns/decorator/beverages"

func AddSoy(b beverages.Beverage) beverages.Beverage {
	return &condiment{
		beverage:    b,
		cost:        15,
		description: "Soy"}
}
