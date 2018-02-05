package main

import observer "github.com/cosaques/patterns/observer/src"

func main() {
	weatherData := new(observer.WeatherData)
	currentConditionsDisplay := new(observer.CurrentConditionsDisplay)
	currentConditionsDisplay.Display()

	weatherData.RegisterObserver(currentConditionsDisplay)
	weatherData.SetMeasurements(32, 75, 500)

	weatherData.RemoveObserver(currentConditionsDisplay)
	weatherData.SetMeasurements(35, 74, 400)

	weatherData.RegisterObserver(currentConditionsDisplay)
	weatherData.RegisterObserver(currentConditionsDisplay)
	weatherData.SetMeasurements(36, 73, 300)
}
