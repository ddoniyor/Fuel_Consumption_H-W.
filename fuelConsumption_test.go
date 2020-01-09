package main

import "testing"

func Test_approximateDistance(t *testing.T) {

	tests := []struct {
		name              string
		distance          int
		consumptionOfFuel int
		amountOfFuel      int
		want              int
	}{
		// TODO: Add test cases.
		{"Approximate Distance", 100, 6, 3, 45},
	}
	for _, test := range tests {
		got := approxDistance(test.distance, test.consumptionOfFuel, test.amountOfFuel)
		if got != test.want {
			t.Error("Fuel test", test.name, "got:", got, "want:", test.want)
		}

	}

}
