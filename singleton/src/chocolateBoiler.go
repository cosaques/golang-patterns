package chocolate

import (
	"fmt"
)

type chocolateBoiler struct {
	empty  bool
	boiled bool
}

func (b *chocolateBoiler) Fill() {
	if b.empty {
		b.empty = false
		b.boiled = false
		fmt.Println("Filling...")
	}
}

func (b *chocolateBoiler) Boil() {
	if !b.empty && !b.boiled {
		b.boiled = true
		fmt.Println("Boiling...")
	}
}

func (b *chocolateBoiler) Drain() {
	if !b.empty && b.boiled {
		b.empty = true
		b.boiled = false
		fmt.Println("Draining...")
	}
}

// Singleton implementation
var boiler *chocolateBoiler
var ch = make(chan bool, 1)

func GetChocolateBoiler() *chocolateBoiler {
	if boiler == nil {
		ch <- true
		if boiler == nil {
			boiler = &chocolateBoiler{empty: true, boiled: false}
		}
		<-ch
	}
	return boiler
}
