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

func TestBoothFunction(t *testing.T) {
	booth := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return ((x + (2 * y) - 7) * (x + (2 * y) - 7)) + (((2 * x) + y - 5) * ((2 * x) + y - 5))
	}

	bounds := []interval.Interval{*interval.New(-10, 10), *interval.New(-10, 10)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(booth)

	minArgs := []float64{1, 3}
	shouldBe := booth(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Booth function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestBukinFunction(t *testing.T) {
	bukin := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return 100*math.Sqrt(math.Abs(y-(0.01*x*x))) + (0.01 * math.Abs(x+10))
	}

	bounds := []interval.Interval{*interval.New(-15, -5), *interval.New(-3, 3)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(bukin)

	minArgs := []float64{-10, 1}
	shouldBe := bukin(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Bukin function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestMatyasFunction(t *testing.T) {
	matyas := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return 0.26*((x*x)+(y*y)) - (0.48 * x * y)
	}

	bounds := []interval.Interval{*interval.New(-15, -5), *interval.New(-3, 3)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(matyas)

	minArgs := []float64{0, 0}
	shouldBe := matyas(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Matyas function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestLeviFunction(t *testing.T) {
	levi := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return (math.Sin(3*math.Pi*x) * math.Sin(3*math.Pi*x)) + (((x - 1) * (x - 1)) * (1 + (math.Sin(3*math.Pi*y) * math.Sin(3*math.Pi*y)))) + ((y - 1) * (y - 1) * (1 + (math.Sin(3*math.Pi*y) * math.Sin(3*math.Pi*y))))
	}

	bounds := []interval.Interval{*interval.New(-10, 10), *interval.New(-10, 10)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(levi)

	minArgs := []float64{1, 1}
	shouldBe := levi(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Lévi function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestEggholderFunction(t *testing.T) {
	eggholder := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return -((y + 47) * math.Sin(math.Sqrt(math.Abs(y+(x/2)+47)))) - (x * math.Sin(math.Sqrt(math.Abs(x-(y+47)))))
	}

	bounds := []interval.Interval{*interval.New(-512, 512), *interval.New(-512, 512)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(eggholder)

	minArgs := []float64{512, 404.2319}
	shouldBe := eggholder(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Eggholder function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestMcCormickFunction(t *testing.T) {
	mcCormick := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return math.Sin(x+y) + ((x - y) * (x - y)) - (1.5 * x) + (2.5 * y) + 1
	}

	bounds := []interval.Interval{*interval.New(-1.5, 4), *interval.New(-3, 4)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(mcCormick)

	minArgs := []float64{-0.54719, -1.54719}
	shouldBe := mcCormick(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("McCormick function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}

func TestSchaffer2Function(t *testing.T) {
	schaffer2 := func(args []float64) float64 {
		x := args[0]
		y := args[1]
		return 0.5 + (((math.Sin((x*x)-(y*y)) * math.Sin((x*x)-(y*y))) - 0.5) / ((1 + (1e-3*(x*x) - (y * y))) * (1 + (1e-3*(x*x) - (y * y)))))
	}

	bounds := []interval.Interval{*interval.New(-100, 100), *interval.New(-100, 100)}

	gwo := New(bounds, 200, 200)
	min := gwo.Min(schaffer2)

	minArgs := []float64{0, 0}
	shouldBe := schaffer2(minArgs)

	if min > shouldBe+error && min < shouldBe-error {
		t.Errorf("Schaffer N.2. function min error! should be: %f with error:+/-%f but found: %f", shouldBe, error, min)
	}
}
