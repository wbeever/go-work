package main

import (
	"fmt"
)

const uSixteenBitMax float64 = 65535
const kmhMultiplE float64 = 1.60934

type car1 struct {
	gasPedal      uint16 // min 0 max 65535
	brakePedal    uint16
	steeringWheel int16 // -32k - +32k
	topSpeedKmh   float64
}

func (c car1) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixteenBitMax)
}
func (c *car1) mph() float64 {
	c.topSpeedKmh = 500
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixteenBitMax / kmhMultiplE)
}

func newerTopSpeed(c car1, speed float64) car1 {
	c.topSpeedKmh = speed
	return c
}

func (c *car1) newTopSpeed(newspeed float64) {
	c.topSpeedKmh = newspeed
}
func main() {
	aCar := car1{
		gasPedal:      65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0,
	}
	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	//aCar.newTopSpeed(500)
	aCar = newerTopSpeed(aCar, 500)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
