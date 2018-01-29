package main

import (
	"github.com/cosaques/ducks"
)

func main() {
	mallardDuck := ducks.NewMallardDuck()
	mallardDuck.Display()
	mallardDuck.PerformFly()

	rubberDuck := ducks.NewRubberDuck()
	rubberDuck.Display()
	rubberDuck.PerformFly()
}
