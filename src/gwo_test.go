package gwo

import (
	"math"
	"testing"
)

const error float64 = 1e-3

func TestEasomFunction(t *testing.T) {
	eosom := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return -math.Cos(x) * math.Cos(y) * math.Exp(-(((x - math.Pi) * (x - math.Pi)) + ((y - math.Pi) * (y - math.Pi))))
	}

	bounds := []Interval{*NewInterval(-100, 100), *NewInterval(-100, 100)}

	gow := NewGWO(bounds, 200, 200)
	min := gow.Min(eosom)

	minArgs := []float64{math.Pi, math.Pi}
	shouldBe := eosom(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestAckleysFunction(t *testing.T) {
	ackleys := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return -20*math.Exp(-0.2*math.Sqrt(0.5*((x*x)+(y*y)))) - math.Exp(0.5*(math.Cos(2*math.Pi*x)+math.Cos(2*math.Pi*y))) + 20 + math.E
	}

	bounds := []Interval{*NewInterval(-5, 5), *NewInterval(-5, 5)}

	gow := NewGWO(bounds, 200, 200)
	min := gow.Min(ackleys)

	minArgs := []float64{0, 0}
	shouldBe := ackleys(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}
