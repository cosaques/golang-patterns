package main

import (
	"fmt"
	"time"

	"github.com/cosaques/patterns/singleton/src"
)

func main() {
	var boiler, boiler2 interface{}
	go func() { boiler = chocolate.GetChocolateBoiler() }()
	go func() { boiler2 = chocolate.GetChocolateBoiler() }()
	time.Sleep(1 * time.Second)
	fmt.Println(&boiler, &boiler2, boiler == boiler2)
	// boiler.Fill()
	// boiler.Boil()
	// boiler.Drain()
}
