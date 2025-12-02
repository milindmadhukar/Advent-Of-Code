package utils

import "golang.org/x/exp/constraints"

func GeneratePermutations[T any](src []T, size int) func(yield func([]T) bool) {
	var res = make([]T, size)
	var generate func(int) bool

	return func(yield func([]T) bool) {
		generate = func(idx int) bool {
			if idx == size {
				if !yield(res) {
					return false
				}

				return true
			}

			for _, v := range src {
				res[idx] = v
				if !generate(idx + 1) {
					return false
				}
			}
			return true
		}

		generate(0)
	}
}

func GenerateCombinations[T any](src []T, size int) func(yield func([]T) bool) {
	var res = make([]T, size)
	var generate func(int, int) bool

	return func(yield func([]T) bool) {
		generate = func(idx, start int) bool {
			if idx == size {
				if !yield(append([]T{}, res...)) {
					return false
				}
				return true
			}

			for i := start; i < len(src); i++ {
				res[idx] = src[i]
				if !generate(idx+1, i+1) {
					return false
				}
			}
			return true
		}

		generate(0, 0)
	}
}

func Permutations[T any](src []T, size int) [][]T {
	var result [][]T
	gen := GeneratePermutations(src, size)
	gen(func(perm []T) bool {
		result = append(result, append([]T{}, perm...))
		return true
	})
	return result
}

func Combinations[T any](src []T, size int) [][]T {
	var result [][]T
	gen := GenerateCombinations(src, size)
	gen(func(comb []T) bool {
		result = append(result, comb)
		return true
	})
	return result
}

func GenerateRange[T constraints.Integer](vals ...T) func(yield func(T) bool) {
	var start, end, step T
	if len(vals) == 0 {
		panic("GenerateRange requires at least one argument")
	}
	if len(vals) == 1 {
		start = T(0)
		end = vals[0]
		step = T(1)
	} else if len(vals) == 2 {
		start = vals[0]
		end = vals[1]
		step = T(1)
	} else if len(vals) == 3 {
		start = vals[0]
		end = vals[1]
		step = vals[2]
	} else {
		panic("GenerateRange requires at most three arguments")
	}

	return func(yield func(T) bool) {
		if step > 0 {
			for i := start; i < end; i += step {
				if !yield(i) {
					return
				}
			}
		} else if step < 0 {
			for i := start; i > end; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}
}
