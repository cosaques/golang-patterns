package condiments

import decorator "github.com/cosaques/patterns/decorator/src"

type Soy struct {
	Beverage decorator.Beverage
}

func (c *Soy) Cost() int64 {
	return c.Beverage.Cost() + 15
}

func (c *Soy) GetDescription() string {
	return c.Beverage.GetDescription() + ", Soy"
}
