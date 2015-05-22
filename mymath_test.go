package mymath

import (
	"fmt"
	"testing"
)

func TestSqurt(t *testing.T) {
	actualNums := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expectedNums := []float64{
		1.0,
		1.4142135,
		1.7320508,
		2.0,
		2.2360679,
	}

	for i, exp := range expectedNums {
		actual := fmt.Sprintf("%1.8f", Squrt(actualNums[i]))
		expected := fmt.Sprintf("%1.7f", exp)
		if actual[:9] != expected {
			t.Errorf("got %v \nwant %v", actual, expected)
		}
	}
}
