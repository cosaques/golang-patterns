package beverages

import "fmt"

type tea struct {
	*abstractBeverage
}

func (b *tea) brew() {
	fmt.Println("Brewing tea...")
}

func (b *tea) addCondiments() {
	fmt.Println("Adding lemon...")
}

func NewTea() *tea {
	result := tea{}
	result.abstractBeverage = &abstractBeverage{&result}

	return &result
}
