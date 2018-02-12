package chocolate

import "fmt"

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

func (b *chocolateBoiler) Drain() {
	if !b.empty && b.boiled {
		b.empty = true
		fmt.Println("Draining...")
	}
}

func (b *chocolateBoiler) Boil() {
	if !b.empty && !b.boiled {
		b.boiled = true
		fmt.Println("Boiling...")
	}
}
