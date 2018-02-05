package main

import strategy "github.com/cosaques/patterns/strategy/src"

func main() {
	mallardDuck := strategy.NewMallardDuck()
	mallardDuck.Display()
	mallardDuck.PerformFly()

	rubberDuck := strategy.NewRubberDuck()
	rubberDuck.Display()
	rubberDuck.PerformFly()
}
