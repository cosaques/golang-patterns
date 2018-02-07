package decorator

type Beverage interface {
	GetDescription() string
	Cost() int64
}
