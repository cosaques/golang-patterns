package condiments

import decorator "github.com/cosaques/patterns/decorator/src"

type Mocha struct {
	Beverage decorator.Beverage
}

func (c *Mocha) Cost() int64 {
	return c.Beverage.Cost() + 20
}

func (c *Mocha) GetDescription() string {
	return c.Beverage.GetDescription() + ", Mocha"
}
