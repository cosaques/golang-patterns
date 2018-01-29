package main

import "github.com/cosaques/ducks"

func main() {
	mallard_duck := ducks.MallardDuck{ducks.Duck{"Me"}}
	mallard_duck.Quack()
	mallard_duck.Display()
}
