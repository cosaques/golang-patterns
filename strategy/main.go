package main

func main() {
	mallardDuck := NewMallardDuck()
	mallardDuck.Display()
	mallardDuck.PerformFly()

	rubberDuck := NewRubberDuck()
	rubberDuck.Display()
	rubberDuck.PerformFly()
}
