package commands

import "fmt"

type Empty struct{}

func (c *Empty) Execute() {
	fmt.Println("No command set")
}
