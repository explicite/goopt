package interval

import (
	"math/rand"
	"time"
)

type Interval struct {
	min, max float64
	rand     *rand.Rand
}

func New(min, max float64) *Interval {
	if min > max {
		return nil
	}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	interval := Interval{min, max, rand}
	return &interval
}

func (interval *Interval) NextRand() float64 {
	return ((interval.rand.Float64() * (interval.max - interval.min)) + float64(1)) + interval.min - float64(1)
}

func (interval *Interval) IsIn(i float64) bool {
	if i >= interval.min && i <= interval.max {
		return true
	} else {
		return false
	}
}
