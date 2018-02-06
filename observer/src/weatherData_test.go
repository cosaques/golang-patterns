package observer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type observerMock struct {
	temperature float32
	humidity    float32
	pressure    float32
}

func (o *observerMock) Update(temperature float32, humidity float32, pressure float32) {
	o.humidity = humidity
	o.temperature = temperature
	o.pressure = pressure
}

func TestRegisterObserver_ShouldBeUpdatedWithCorrectParameters(t *testing.T) {
	weatherData := new(WeatherData)
	observer := new(observerMock)

	weatherData.RegisterObserver(observer)

	temperature := float32(1)
	humidity := float32(2)
	pressure := float32(3)
	weatherData.SetMeasurements(temperature, humidity, pressure)

	assert.Equal(t, temperature, observer.temperature, "Temperature should be updated")
	assert.Equal(t, humidity, observer.humidity, "Humidity should be updated")
	assert.Equal(t, pressure, observer.pressure, "Pressure should be updated")
}

func TestRemoveObserver_ShouldNotBeUpdatedAfterBeRomoved(t *testing.T) {
	weatherData := new(WeatherData)
	observer := new(observerMock)

	weatherData.RegisterObserver(observer)
	weatherData.RemoveObserver(observer)

	temperature := float32(1)
	humidity := float32(2)
	pressure := float32(3)
	weatherData.SetMeasurements(temperature, humidity, pressure)

	assert.NotEqual(t, temperature, observer.temperature, "Temperature should not be updated")
	assert.NotEqual(t, humidity, observer.humidity, "Humidity should not be updated")
	assert.NotEqual(t, pressure, observer.pressure, "Pressure should not be updated")
}
