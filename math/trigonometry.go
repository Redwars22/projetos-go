package main

import (
	"fmt"
	"math"
)

func convertToRadians(angle float64) float64 {
	return ((angle * math.Pi) / 180)
}

func main() {
	angle := 120.0
	rad := convertToRadians(angle)

	fmt.Println("CÁLCULO DE SEN, COS, TG COM GO")

	fmt.Println("SENO DE " + fmt.Sprintf("%v", angle) + "º: " + fmt.Sprintf("%v", math.Sin(rad)))
	fmt.Println("COSSENO DE " + fmt.Sprintf("%v", angle) + "º: " + fmt.Sprintf("%v", math.Cos(rad)))
	fmt.Println("TANGENTE DE " + fmt.Sprintf("%v", angle) + "º: " + fmt.Sprintf("%v", math.Tan(rad)))
}
