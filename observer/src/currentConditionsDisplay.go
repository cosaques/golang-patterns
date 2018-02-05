package observer

import (
	"fmt"
)

type CurrentConditionsDisplay struct {
	temperature float32
	humidity    float32
}

func (c *CurrentConditionsDisplay) Update(temperature float32, humidity float32, pressure float32) {
	c.temperature = temperature
	c.humidity = humidity
	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions : %fF degrees and humidity of %f\n", c.temperature, c.humidity)
}
