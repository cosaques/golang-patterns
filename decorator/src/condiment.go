package decorator

type Condiment interface {
	GetDescription() string
	Cost() int64
}
