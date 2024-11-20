package space

type Planet string

const EOP = 31557600 // Earth's orbital period in seconds

var ageFactors = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	if factor, ok := ageFactors[planet]; ok {
		return seconds / EOP / factor
	} else {
		return -1 // or handle error more appropriately
	}
}
