package interval

import "testing"

func TestNextRand(t *testing.T) {
	interval := New(1, 3)
	var rand float64
	for i := 0; i <= 1000; i++ {
		rand = interval.NextRand()
		if rand > 3 || rand < 1 {
			t.Errorf("rand:%f out of the bounds [%f,%f]", rand, interval.min, interval.max)
		}
	}
}
