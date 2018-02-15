package commands

import (
	"fmt"
)

type LightOn struct{}

func (c *LightOn) Execute() {
	fmt.Println("Light is on")
}
