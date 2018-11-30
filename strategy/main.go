package main

import strategy "github.com/cosaques/patterns/strategy/src"

func main() {
	var duck strategy.Duck

	duck = strategy.NewMallardDuck()
	duck.Display()
	duck.Swim()
	duck.PerformFly()

	duck = strategy.NewRubberDuck()
	duck.Display()
	duck.Swim()
	duck.PerformFly()
}
