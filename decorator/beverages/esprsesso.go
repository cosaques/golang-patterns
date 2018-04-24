package beverages

type Espresso struct {
}

func (e *Espresso) GetDescription() string {
	return "Hot Espresso"
}

func (e *Espresso) Cost() int64 {
	return 199
}
