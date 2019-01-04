package beverages

import "fmt"

type tea struct {
	*baseBeverage
}

func (b *tea) brew() {
	fmt.Println("Brewing tea...")
}

func (b *tea) addCondiments() {
	fmt.Println("Adding lemon...")
}

func NewTea() *tea {
	result := tea{}
	result.baseBeverage = &baseBeverage{&result}

	return &result
}
