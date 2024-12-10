package main

import (
	"testing"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/solutions"
)

func BenchmarkDay05(b *testing.B) {
	sol := solutions.GetSolution(5)
	for i := 0; i < b.N; i++ {
		sol.Part1()
		sol.Part2()
	}
}

func BenchmarkDay09(b *testing.B) {
	sol := solutions.GetSolution(9)
	for i := 0; i < b.N; i++ {
		sol.Part1()
		sol.Part2()
	}
}

// go test -bench=BenchmarkDay05 -cpuprofile profiles/cpu.pprof . 
