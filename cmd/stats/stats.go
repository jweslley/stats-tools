package main

import (
	"fmt"
	"github.com/jweslley/stats-tools"
)

func main() {
	stats.Tool("Output a summary table including mean, median, mininum, maximum, standard deviation, variance and number count of a number sequence", func(s *stats.Stats) {
		fmt.Printf("Min: %.6f\nMean: %.6f\nMax: %.6f\nStdDev: %.6f\nVar: %.6f\nCount: %d\n",
			s.Min(), s.Mean(), s.Max(), s.StdDev(), s.Variance(), s.Count())
	})
}
