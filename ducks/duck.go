package ducks

import (
	"fmt"
)

// Duck is a base class for all the ducks
type Duck struct {
	Name string
}

// Display shows the greeting info of the duck
func (d *Duck) Display() {
	fmt.Printf("Hello my name is %s", d.Name)
}
