package condiments

import (
	"github.com/cosaques/patterns/decorator/src"
)

func AddSoy(b decorator.Beverage) decorator.Beverage {
	return &condiment{
		beverage:    b,
		cost:        15,
		description: "Soy"}
}
