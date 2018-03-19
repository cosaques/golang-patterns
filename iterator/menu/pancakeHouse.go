package menu

type pancakeHouse struct {
	items map[string]*MenuItem
}

func (m *pancakeHouse) GetItems() map[string]*MenuItem {
	return m.items
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
