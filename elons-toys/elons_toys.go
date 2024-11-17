package elon

import "fmt"

// Drive updates the car's distance and battery based on its speed and battery drain
func (c *Car) Drive() {
    // Early return if not enough battery
    if c.battery < c.batteryDrain {
        return
    }
    
    // Update battery and distance in one operation
    c.battery -= c.batteryDrain
    c.distance += c.speed
}

// DisplayDistance returns the current distance as a formatted string
func (c *Car) DisplayDistance() string {
    // Using const for string format to improve maintainability
    const distanceFormat = "Driven %d meters"
    return fmt.Sprintf(distanceFormat, c.distance)
}

// DisplayBattery returns the current battery level as a formatted string
func (c *Car) DisplayBattery() string {
    // Using const for string format to improve maintainability
    const batteryFormat = "Battery at %d%%"
    return fmt.Sprintf(batteryFormat, c.battery)
}

// CanFinish determines if the car can complete the track with current battery
func (c *Car) CanFinish(trackDistance int) bool {
    if trackDistance <= 0 {
        return true
    }
    
    // Calculate maximum possible distance with current battery
    remainingDrives := c.battery / c.batteryDrain
    maxDistance := remainingDrives * c.speed
    
    return maxDistance >= trackDistance
}
