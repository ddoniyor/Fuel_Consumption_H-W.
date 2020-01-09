package main

import "testing"

func Test_approximateDistance(t *testing.T) {

	tests := []struct {
		name           string
		distOfConsFuel int
		fuelConsume    int
		fuelAmount     int
		want           int
	}{

		{"Approximate Distance", 100, 6, 3, 45},
	}
	for _, test := range tests {
		got := approxDistance(test.distOfConsFuel, test.fuelConsume, test.fuelAmount)
		if got != test.want {
			t.Error("Fuel test", test.name, "got:", got, "want:", test.want)
		}
	}
}
