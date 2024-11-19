package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius TemperatureUnit = iota
	Fahrenheit
	Kelvin
)

func (t TemperatureUnit) String() string {
	switch t {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	case Kelvin:
		return "K"
	default:
		return "Unknown"
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour SpeedUnit = iota
	MilesPerHour
)

func (s SpeedUnit) String() string {
	switch s {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		return "Unknown"
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
	return fmt.Sprintf(
		"%s: %s, Wind %s at %s, %d%% Humidity",
		m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity,
	)
}
