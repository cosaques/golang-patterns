package restaurant

import (
	"github.com/cosaques/patterns/iterator/iterator"
)

type pancakeMenu struct {
	items map[string]*MenuItem
}

type pancakeMenuIterator struct {
	items    []interface{}
	position int
}

// pancakeMenu methods

func (m *pancakeMenu) CreateIterator() iterator.Iterator {
	menuItems := make(map[string]interface{})
	for k, v := range m.items {
		menuItems[k] = v
	}
	return newPancakeMenuIterator(menuItems)
}

func (m *pancakeMenu) addItem(name string, description string, vegeterian bool, price float32) {
	m.items[name] = &MenuItem{name, description, vegeterian, price}
}

func NewPancakeMenu() *pancakeMenu {
	result := new(pancakeMenu)
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

// pancakeMenuIterator methods

func (i *pancakeMenuIterator) Next() interface{} {
	next := i.items[i.position]
	i.position += 1
	return next
}

func (i *pancakeMenuIterator) HasNext() bool {
	return len(i.items)-1 >= i.position
}

func newPancakeMenuIterator(m map[string]interface{}) *pancakeMenuIterator {
	result := &pancakeMenuIterator{}
	for _, v := range m {
		result.items = append(result.items, v)
	}

	return result
}
