package utils

func Permutations[T any](src []T, size int) func(yield func([]T) bool) {
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

func Combinations[T any](src []T, size int) func(yield func([]T) bool) {
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
