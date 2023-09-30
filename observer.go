package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Observer defines the observer interface.
type Observer interface {
	Update(temperature, humidity, pressure float64)
}


type DisplayElement interface {
	Display()
}

// WeatherData represents the subject that tracks weather measurements.
type WeatherData struct {
	observers    []Observer
	temperature  float64
	humidity     float64
	pressure     float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) RegisterObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherData) RemoveObserver(observer Observer) {
	for i, o := range w.observers {
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	
	w.temperature = rand.Float64()*100 - 50
	w.humidity = rand.Float64() * 100
	w.pressure = rand.Float64() * 50

	w.NotifyObservers()
}

// CurrentConditionsDisplay represents a concrete observer that displays current weather conditions.
type CurrentConditionsDisplay struct {
	temperature  float64
	humidity     float64
	weatherData  *WeatherData
}

func NewCurrentConditionsDisplay(weatherData *WeatherData) *CurrentConditionsDisplay {
	display := &CurrentConditionsDisplay{
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(display)
	return display
}

func (ccd *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2f°F and %.2f%% humidity\n", ccd.temperature, ccd.humidity)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Creating a WeatherData object as the subject.
	weatherData := NewWeatherData()

	// Creating current conditions display and registering it with the subject.
	currentDisplay := NewCurrentConditionsDisplay(weatherData)

	// Simulating weather data changes and notifying observers.
	for i := 0; i < 5; i++ {
		weatherData.MeasurementsChanged()
		time.Sleep(2 * time.Second)
	}
}
