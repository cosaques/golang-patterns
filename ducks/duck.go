package ducks

import (
	"fmt"
)

// Duck is a base class for all the ducks
// http://blog.bingecoder.net/index.php/2016/06/20/abstract-classes-in-golang/
type Duck struct {
	Name string
}

// Display shows the greeting info of the duck
func (d *Duck) Display() {
	panic("Abstract method")
}

func (d *Duck) Quack() {
	fmt.Println("Kwaaaaa")
}

type MallardDuck struct {
	Duck
}

func (d *MallardDuck) Display() {
	fmt.Printf("Hello my name is %s", d.Name)
}
