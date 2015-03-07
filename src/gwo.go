package gwo

import "math"

type Function func([]float64) float64

type Wolf struct {
	coordinates []float64
}

func (wolf *Wolf) ToSpace(intervals *[]Interval) {
	for i := range wolf.coordinates {
		if !(*intervals)[i].IsIn(wolf.coordinates[i]) {
			wolf.coordinates[i] = (*intervals)[i].NextRand()
		}
	}
}

type GWO struct {
	bounds     []Interval
	iterations uint
	wolfs      []Wolf
}

func NewGWO(bounds []Interval, noWolfs uint, noIterations uint) *GWO {
	wolfs := make([]Wolf, noWolfs)

	for i := range wolfs {
		coordinates := make([]float64, len(bounds))
		for c := range coordinates {
			coordinates[c] = bounds[c].NextRand()
		}
		wolfs[i].coordinates = coordinates
		wolfs[i].ToSpace(&bounds)
	}

	gwo := GWO{bounds, noIterations, wolfs}
	return &gwo
}

func (gwo *GWO) Min(function Function) float64 {
	return gwo.optimize(function, math.Inf(0))
}

func (gwo *GWO) Max(function Function) float64 {
	return gwo.optimize(function, math.Inf(-1))
}

func (gwo *GWO) optimize(function Function, inf float64) float64 {
	iterations := gwo.iterations

	alphaWolf := Wolf{make([]float64, len(gwo.bounds))}
	betaWolf := Wolf{make([]float64, len(gwo.bounds))}
	deltaWolf := Wolf{make([]float64, len(gwo.bounds))}

	alphaScore := inf
	betaScore := inf
	deltaScore := inf

	for iteration := uint(0); iteration < iterations; iteration++ {
		for i := range gwo.wolfs {
			wolf := &gwo.wolfs[i]
			wolf.ToSpace(&gwo.bounds)

			fitness := function(wolf.coordinates)

			if fitness < alphaScore {
				alphaScore = fitness
				alphaWolf = *wolf
			}

			if fitness > alphaScore && fitness < betaScore {
				betaScore = fitness
				betaWolf = *wolf
			}

			if fitness > alphaScore && fitness > betaScore && fitness < deltaScore {
				deltaScore = fitness
				deltaWolf = *wolf
			}
		}

		coefficient := 2.0 - float64(iteration)*(2.0/float64(iterations))

		for i := range gwo.wolfs {
			wolf := &gwo.wolfs[i]
			for j := range wolf.coordinates {
				rand := gwo.bounds[j].rand

				aX, aY := 2.0*coefficient*rand.Float64()-1.0, 2.0*rand.Float64()
				dAlpha := math.Abs(aY*alphaWolf.coordinates[j] - wolf.coordinates[j])
				x1 := alphaWolf.coordinates[j] - aX*dAlpha

				bX, bY := 2.0*coefficient*rand.Float64()-1.0, 2.0*rand.Float64()
				dBeta := math.Abs(bY*betaWolf.coordinates[j] - wolf.coordinates[j])
				x2 := betaWolf.coordinates[j] - bX*dBeta

				dX, dY := 2.0*coefficient*rand.Float64()-1.0, 2.0*rand.Float64()
				dDelta := math.Abs(dY*deltaWolf.coordinates[j] - wolf.coordinates[j])
				x3 := deltaWolf.coordinates[j] - dX*dDelta

				wolf.coordinates[j] = (x1 + x2 + x3) / 3.0
			}
		}
	}

	return alphaScore
}
