package main

import (
	"github.com/cosaques/patterns/singleton/src"
)

func main() {
	boiler := chocolate.GetChocolateBoiler()
	boiler.Fill()
	boiler.Boil()
	boiler.Drain()
}
