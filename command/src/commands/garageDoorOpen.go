package commands

import (
	"fmt"
)

type GarageDoorOpen struct{}

func (c *GarageDoorOpen) Execute() {
	fmt.Println("Garage door is opened")
}
