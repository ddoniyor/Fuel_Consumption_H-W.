package main

import "fmt"

func main() {
	const distance = 100
	fuelConsume := 6
	fuelAmount := 3

	result := approxDistance(distance, fuelConsume, fuelAmount)

	fmt.Println(result, "km")
}

func approxDistance(distance int, fuelConsume int, fuelAmount int) int {
	var approxDistance int

	approxDistance = distance * fuelAmount / fuelConsume
	tenPercent := approxDistance / 10

	return approxDistance - tenPercent
}
