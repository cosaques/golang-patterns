package commands

import (
	"fmt"
)

type LightOff struct{}

func (c *LightOff) Execute() {
	fmt.Println("Light is off")
}
