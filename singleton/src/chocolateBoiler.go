package chocolate

import (
	"fmt"
	"sync"
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
var mux sync.Mutex

func GetChocolateBoiler() *chocolateBoiler {
	if boiler == nil {
		mux.Lock()
		if boiler == nil {
			boiler = &chocolateBoiler{empty: true, boiled: false}
		}
		mux.Unlock()
	}
	return boiler
}
