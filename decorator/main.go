package main

import (
	"fmt"

	"github.com/cosaques/patterns/decorator/src"
	"github.com/cosaques/patterns/decorator/src/beverages"
	"github.com/cosaques/patterns/decorator/src/condiments"
)

func main() {
	var beverage1 decorator.Beverage = new(beverages.Espresso)
	printBeverage(beverage1)

	var beverage2 decorator.Beverage = new(beverages.HouseBlend)
	beverage2 = condiments.AddMocha(beverage2)
	beverage2 = condiments.AddMocha(beverage2)
	printBeverage(beverage2)

	var beverage3 decorator.Beverage = new(beverages.Espresso)
	beverage3 = condiments.AddSoy(beverage3)
	beverage3 = condiments.AddMocha(beverage3)
	printBeverage(beverage3)
}

func printBeverage(b decorator.Beverage) {
	fmt.Printf("%s $%v\n", b.GetDescription(), b.Cost())
}
