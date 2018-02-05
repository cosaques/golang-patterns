package observer

import "github.com/cosaques/patterns/observer/src/abstract"

type WeatherData struct {
	temperature float32
	humidity    float32
	pressure    float32
	observers   []abstract.Observer
}

func (w *WeatherData) RegisterObserver(o abstract.Observer) {
	for _, obs := range w.observers {
		if obs == o {
			return
		}
	}
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o abstract.Observer) {
	for i, obs := range w.observers {
		if obs == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			return
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, obs := range w.observers {
		obs.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) measurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature float32, humidity float32, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}
