package speed

// Car defines the properties of a remote-controlled car.
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote-controlled car with a full battery.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// Track defines the properties of a race track.
type Track struct {
	distance int
}

// NewTrack creates a new race track with the specified distance.
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive moves the car forward once if there is enough battery.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car
}

// CanFinish checks if a car can complete a track.
func CanFinish(car Car, track Track) bool {
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	return maxDistance >= track.distance
}
