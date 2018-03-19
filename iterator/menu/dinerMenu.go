package menu

type dinerMenu struct {
	items []*MenuItem
}

type dinerMenuIterator struct {
	items    []interface{}
	position int
}

// dinerMenu methods

func (m *dinerMenu) GetItems() []*MenuItem {
	return m.items
}

func (m *dinerMenu) addItem(name string, description string, vegeterian bool, price float32) {
	m.items = append(m.items, &MenuItem{name, description, vegeterian, price})
}

func NewDinerMenu() *dinerMenu {
	result := new(dinerMenu)
	result.items = make([]*MenuItem, 6)

	result.addItem("Vegeterian BLT",
		"Soja Becon with lettuce & tomato",
		true,
		12.49)

	result.addItem("BLT",
		"Becon with egg",
		false,
		15.00)

	result.addItem("Hotdog",
		"A hotdog with onions, ketchup",
		false,
		2.00)

	return result
}

// dinerMenuIterator methods

func (i *dinerMenuIterator) Next() interface{} {
	next := i.items[i.position]
	i.position += 1
	return next
}

func (i *dinerMenuIterator) HasNext() bool {
	return len(i.items)-1 >= i.position
}