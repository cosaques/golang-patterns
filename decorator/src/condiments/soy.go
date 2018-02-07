package condiments

import (
	"github.com/cosaques/patterns/decorator/src"
)

type soy struct {
	condiment
}

func AddSoy(b decorator.Beverage) *soy {
	return &soy{condiment{
		beverage:    b,
		cost:        15,
		description: "Soy"}}
}
