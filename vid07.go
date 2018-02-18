package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car2 struct {
	gasPedal      uint16 // min 0 max 65535
	brakePedal    uint16
	steeringWheel int16 // -32k - +32k
	topSpeedKmh   float64
}

func (c car2) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}
func (c car2) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}
func main() {
	aCar := car2{
		gasPedal:      65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0,
	}
	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
