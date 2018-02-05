package observer

import (
	"fmt"

	"github.com/cosaques/patterns/observer/src/abstract"
)

type currentConditionsDisplay struct {
	temperature float32
	humidity    float32
	weatherData abstract.Subject
}

func (c *currentConditionsDisplay) Update(temperature float32, humidity float32, pressure float32) {
	c.temperature = temperature
	c.humidity = humidity
	c.Display()
}

func (c *currentConditionsDisplay) Display() {
	fmt.Printf("Current conditions : %fF degrees and humidity of %f\n", c.temperature, c.humidity)
}

func NewCurrentConditionsDisplay(weatherData abstract.Subject) *currentConditionsDisplay {
	c := new(currentConditionsDisplay)
	c.weatherData = weatherData
	weatherData.RegisterObserver(c)

	return c
}
