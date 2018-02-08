package pizzas

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
