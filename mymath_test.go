package mymath

import (
	"math"
	"testing"
)

func TestSqurt(t *testing.T) {

	var tests = []struct {
		num  float64
		want float64
	}{
		{1.0, 1.0},
		{2.0, 1.4142135},
		{3.0, 1.7320508},
		{4.0, 2.0},
		{5.0, 2.2360679},
	}

	for _, test := range tests {
		got := Squrt(test.num)
		if math.Abs(got-test.want) > 0.0000001 {
			t.Errorf("Squrt(%1.8f) returned %1.8f, want %1.8f", test.num, got, test.want)
		}
	}
}
