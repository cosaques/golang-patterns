package restaurant

import (
	"fmt"

	"github.com/cosaques/patterns/iterator/iterator"
)

type Waitress struct {
	pancakeMenu *pancakeMenu
	dinerMenu   *dinerMenu
}

func (w *Waitress) PrintMenu() {
	fmt.Println("MENU\n----")
	fmt.Println("Breakfast")
	printMenu(w.pancakeMenu.CreateIterator())
	fmt.Println("Diner")
	printMenu(w.dinerMenu.CreateIterator())
}

func printMenu(i iterator.Iterator) {
	for i.HasNext() {
		menuItem := i.Next().(*MenuItem)
		fmt.Printf("\t%s \n\t\tprice: %.2f \n\t\t%s\n",
			menuItem.name, menuItem.price, menuItem.description)
	}
}

func NewWaitress(pancakeMenu *pancakeMenu, dinerMenu *dinerMenu) *Waitress {
	return &Waitress{pancakeMenu, dinerMenu}
}
