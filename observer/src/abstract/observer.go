package abstract

type Observer interface {
	Update(temperature float32, humidity float32, pressure float32)
}
