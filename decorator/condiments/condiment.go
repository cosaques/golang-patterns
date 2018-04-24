package condiments

import (
	"fmt"

	"github.com/cosaques/patterns/decorator/beverages"
)

type condiment struct {
	beverage    beverages.Beverage
	cost        int64
	description string
}

func (c *condiment) Cost() int64 {
	return c.beverage.Cost() + c.cost
}

func (c *condiment) GetDescription() string {
	return fmt.Sprintf("%s, %s", c.beverage.GetDescription(), c.description)
}
