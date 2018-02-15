package main

import (
	"github.com/cosaques/patterns/singleton/src"
)

func main() {
	var boiler = chocolate.GetChocolateBoiler()
	boiler.Fill()
	boiler.Boil()
	boiler.Drain()
}
