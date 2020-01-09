package main

import "fmt"

func main() {
	const distOfConsFuel = 100
	fuelConsume := 6
	fuelAmount := 3

	resDistance := approxDistance(distOfConsFuel, fuelConsume, fuelAmount)

	fmt.Println(resDistance, "km")
}

func approxDistance(distOfConsFuel int, fuelConsume int, fuelAmount int) int {

	approxDistance := distOfConsFuel * fuelAmount / fuelConsume
	tenPercent := approxDistance / 10

	return approxDistance - tenPercent
}
