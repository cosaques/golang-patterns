package menu

import (
	"github.com/cosaques/patterns/iterator/iterator"
)

type pancakeHouse struct {
	items map[string]*MenuItem
}

type pancakeHouseIterator struct {
	items    []interface{}
	position int
}

// pancakeHouse methods

func (m *pancakeHouse) CreateIterator() iterator.Iterator {
	menuItems := make(map[string]interface{})
	for k, v := range m.items {
		menuItems[k] = v
	}
	return newPancakeHouseIterator(menuItems)
}

func (m *pancakeHouse) addItem(name string, description string, vegeterian bool, price float32) {
	m.items[name] = &MenuItem{name, description, vegeterian, price}
}

func NewPancakeHouse() *pancakeHouse {
	result := new(pancakeHouse)
	result.items = make(map[string]*MenuItem)

	result.addItem("K&B's pancakes breakfast",
		"Deliciuos breakfast for a good price",
		false,
		10.99)

	result.addItem("Blueberry pancakes",
		"Pancakes made with fresh blueberries",
		true,
		7.99)

	return result
}

// pancakeHouseIterator methods

func (i *pancakeHouseIterator) Next() interface{} {
	next := i.items[i.position]
	i.position += 1
	return next
}

func (i *pancakeHouseIterator) HasNext() bool {
	return len(i.items)-1 >= i.position
}

func newPancakeHouseIterator(m map[string]interface{}) *pancakeHouseIterator {
	result := &pancakeHouseIterator{}
	for _, v := range m {
		result.items = append(result.items, v)
	}

	return result
}
