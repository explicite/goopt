package gwo

import (
	"../interval"
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

	bounds := []interval.Interval{*interval.New(-100, 100), *interval.New(-100, 100)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(eosom)

	minArgs := []float64{math.Pi, math.Pi}
	shouldBe := eosom(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Easom function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestAckleysFunction(t *testing.T) {
	ackleys := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return -20*math.Exp(-0.2*math.Sqrt(0.5*((x*x)+(y*y)))) - math.Exp(0.5*(math.Cos(2*math.Pi*x)+math.Cos(2*math.Pi*y))) + 20 + math.E
	}

	bounds := []interval.Interval{*interval.New(-5, 5), *interval.New(-5, 5)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(ackleys)

	minArgs := []float64{0, 0}
	shouldBe := ackleys(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Ackleys function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}
