package observer

import (
	"testing"

	"github.com/cosaques/patterns/observer/src/abstract"

	"github.com/stretchr/testify/assert"
)

type observerMock struct {
	isUpdated bool
	subject   abstract.Subject
}

func (o *observerMock) Update() {
	o.isUpdated = true
}

func (o *observerMock) SetSubject(s abstract.Subject) {
	o.subject = s
}

func TestRegisterObserver_ShouldBeUpdatedAfterBeRegistered(t *testing.T) {
	weatherData := new(WeatherData)
	observer := new(observerMock)

	weatherData.RegisterObserver(observer)

	temperature := float32(1)
	humidity := float32(2)
	pressure := float32(3)
	weatherData.SetMeasurements(temperature, humidity, pressure)

	assert.True(t, observer.isUpdated, "Observer should be updated if added")
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

	assert.False(t, observer.isUpdated, "Observer should not be updated if removed")
}
