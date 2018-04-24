package abstract

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	GetTemperature() float32
	GetHumidity() float32
	GetPressure() float32
}
