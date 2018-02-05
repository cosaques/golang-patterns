package strategy

import (
	"fmt"
)

type flyBehavior func()

var simpleFly = func() {
	fmt.Println("Flying.... youpi !!!!")
}

var noFly = func() {
	fmt.Println("Sorry.... I cannot fly !!!!")
}
