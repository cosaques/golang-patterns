package beverages

type Espresso struct {
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) Cost() int64 {
	return 199
}
