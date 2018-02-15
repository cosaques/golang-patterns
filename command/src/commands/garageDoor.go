package commands

import (
	"fmt"
)

type GarageDoor struct{}

func (c *GarageDoor) Execute() {
	fmt.Println("Garage door is opened")
}
