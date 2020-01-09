package main

import "fmt"

func main() {
	distance:=100
	consumptionOfFuel:=6
	amountOfFuel:=3

	result:=approximateDistance(distance,consumptionOfFuel,amountOfFuel)

	fmt.Println(result,"km")
}

func approximateDistance(distance int,consumptionOfFuel int, amountOfFuel int)int{
	var approximateDistance int
	approximateDistance = distance*amountOfFuel/consumptionOfFuel
	return approximateDistance-approximateDistance/10
}
