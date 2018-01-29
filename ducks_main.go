package main

import (
	"github.com/cosaques/ducks"
)

func main() {
	mallardDuck := ducks.NewMallardDuck()
	mallardDuck.Quack()
	mallardDuck.Display()
}
