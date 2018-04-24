package beverages

type Beverage interface {
	GetDescription() string
	Cost() int64
}
