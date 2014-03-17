package stats

import (
	"math"
)

func NewStats() *Stats {
	return &Stats{
		max:      -math.MaxFloat64,
		min:      math.MaxFloat64,
		variance: [2]float64{-1, 0},
	}
}

type Stats struct {
	count         int64
	sum, min, max float64
	variance      [2]float64
}

func (s *Stats) Count() int64 {
	return s.count
}

func (s *Stats) Sum() float64 {
	return s.sum
}

func (s *Stats) Mean() float64 {
	if s.count == 0 {
		return 0
	}
	return s.sum / float64(s.count)
}

func (s *Stats) Min() float64 {
	if s.count == 0 {
		return 0
	}
	return s.min
}

func (s *Stats) Max() float64 {
	if s.count == 0 {
		return 0
	}
	return s.max
}

func (s *Stats) StdDev() float64 {
	return math.Sqrt(s.Variance())
}

func (s *Stats) Variance() float64 {
	if s.count <= 1 {
		return 0
	}
	return s.variance[1] / float64(s.count-1)
}

func (s *Stats) Update(v float64) {
	s.count++
	s.sum += v
	if v < s.min {
		s.min = v
	}
	if v > s.max {
		s.max = v
	}
	if s.variance[0] == -1 {
		s.variance[0] = v
		s.variance[1] = 0.0
	} else {
		v0 := s.variance[0]
		v1 := s.variance[1]
		s.variance[0] = v0 + (v-v0)/float64(s.count)
		s.variance[1] = v1 + (v-v0)*(v-s.variance[0])
	}
}
