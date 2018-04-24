package beverages

type HouseBlend struct {
}

func (e *HouseBlend) GetDescription() string {
	return "House Blend Coffee"
}

func (e *HouseBlend) Cost() int64 {
	return 89
}
