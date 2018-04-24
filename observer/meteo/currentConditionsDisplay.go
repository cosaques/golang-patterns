package observer

import (
	"fmt"

	"github.com/cosaques/patterns/observer/meteo/abstract"
)

type CurrentConditionsDisplay struct {
	temperature float32
	humidity    float32
	subject     abstract.Subject
}

func (c *CurrentConditionsDisplay) Update() {
	c.temperature = c.subject.GetTemperature()
	c.humidity = c.subject.GetHumidity()
	c.Display()
}

func (c *CurrentConditionsDisplay) SetSubject(s abstract.Subject) {
	c.subject = s
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions : %fF degrees and humidity of %f\n", c.temperature, c.humidity)
}
