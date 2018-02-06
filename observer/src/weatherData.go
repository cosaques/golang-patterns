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
	o.SetSubject(w)
}

func (w *WeatherData) RemoveObserver(o abstract.Observer) {
	for i, obs := range w.observers {
		if obs == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			o.SetSubject(nil)
			return
		}
	}
}

func (w *WeatherData) GetTemperature() float32 {
	return w.temperature
}

func (w *WeatherData) GetHumidity() float32 {
	return w.humidity
}

func (w *WeatherData) GetPressure() float32 {
	return w.pressure
}

func (w *WeatherData) notifyObservers() {
	for _, obs := range w.observers {
		obs.Update()
	}
}

func (w *WeatherData) SetMeasurements(temperature float32, humidity float32, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.notifyObservers()
}
