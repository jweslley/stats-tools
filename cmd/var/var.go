package main

import (
	"fmt"
	"github.com/jweslley/stats-tools"
)

func main() {
	stats.Tool("Calculate the variance of a number sequence", func(s *stats.Stats) {
		fmt.Println(s.Variance())
	})
}
