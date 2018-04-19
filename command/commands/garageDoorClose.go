package commands

import (
	"fmt"
)

type GarageDoorClose struct{}

func (c *GarageDoorClose) Execute() {
	fmt.Println("Garage door is closed")
}
