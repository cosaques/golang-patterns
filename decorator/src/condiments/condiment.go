package condiments

import (
	"fmt"

	decorator "github.com/cosaques/patterns/decorator/src"
)

type condiment struct {
	beverage    decorator.Beverage
	cost        int64
	description string
}

func (c *condiment) Cost() int64 {
	return c.beverage.Cost() + c.cost
}

func (c *condiment) GetDescription() string {
	return fmt.Sprintf("%s, %s", c.beverage.GetDescription(), c.description)
}
