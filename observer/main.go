package main

import observer "github.com/cosaques/patterns/observer/src"

func main() {
	weatherData := new(observer.WeatherData)
	currentConditionsDisplay := observer.NewCurrentConditionsDisplay(weatherData)

	currentConditionsDisplay.Display()
	weatherData.SetMeasurements(32, 75, 500)
}
