package beverages

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coffee struct {
	*baseBeverage
}

func (b *coffee) brew() {
	fmt.Println("Brewing coffee...")
}

func (b *coffee) addCondiments() {
	fmt.Println("Adding milk and sugar...")
}

func (b *coffee) customerWantsCondiments() bool {
	answer := getUserInput()
	if strings.ToLower(answer)[0] == 'y' {
		return true
	}

	return false
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want milk and sugar (y/n): ")
	text, _ := reader.ReadString('\n')

	return text
}

func NewCoffee() *coffee {
	c := &coffee{}
	c.baseBeverage = &baseBeverage{c}
	return c
}
