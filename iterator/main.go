package main

import "github.com/cosaques/patterns/iterator/restaurant"

func main() {
	waitress := restaurant.NewWaitress(restaurant.NewPancakeMenu(), restaurant.NewDinerMenu())
	waitress.PrintMenu()
}
