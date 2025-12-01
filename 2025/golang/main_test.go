package main

import (
	"testing"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/day01"
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/models"
)

func BenchmarkDay01(b *testing.B) {
	var sol models.Solution
	for i := 0; i < b.N; i++ {
		sol = day01.Solve()
		sol.Part1()
		sol.Part2()
	}
}


// go test -bench=BenchmarkDay05 -cpuprofile profiles/cpu.pprof .
