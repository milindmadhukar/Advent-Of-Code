package main

import (
	"testing"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day05"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day06"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day09"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day11"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day19"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/models"
)

func BenchmarkDay05(b *testing.B) {
	var sol models.Solution
	for i := 0; i < b.N; i++ {
		sol = day05.Solve()
		sol.Part1()
		sol.Part2()
	}
}

func BenchmarkDay06(b *testing.B) {
	var sol models.Solution
	for i := 0; i < b.N; i++ {
		sol = day06.Solve()
		sol.Part1()
		sol.Part2()
	}
}

func BenchmarkDay09(b *testing.B) {
	var sol models.Solution
	for i := 0; i < b.N; i++ {
		sol = day09.Solve()
		sol.Part1()
		sol.Part2()
	}
}

func BenchmarkDay11(b *testing.B) {
	var sol models.Solution
	for i := 0; i < b.N; i++ {
		sol = day11.Solve()
		sol.Part1()
		sol.Part2()
	}
}

func BenchmarkDay19(b *testing.B) {
	var sol models.Solution
	for i := 0; i < b.N; i++ {
		sol = day19.Solve()
		sol.Part1()
		sol.Part2()
	}
}

// go test -bench=BenchmarkDay05 -cpuprofile profiles/cpu.pprof .
