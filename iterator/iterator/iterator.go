package iterator

type Iterator interface {
	Next() interface{}
	HasNext() bool
}
