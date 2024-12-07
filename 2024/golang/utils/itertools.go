package utils

func Permutations[T comparable](input []T, size int) [][]T {
	var result [][]T
	current := make([]T, size)
	permutationsHelper(input, size, 0, current, &result)

	return result
}

func permutationsHelper[T comparable](input []T, size int, position int, current []T, result *[][]T) {
	if position == size {
		combination := make([]T, size)
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	for i := 0; i < len(input); i++ {
		current[position] = input[i]
		permutationsHelper(input, size, position+1, current, result)
	}
}

func Combinations[T comparable](input []T, r int) [][]T {
	var result [][]T
	current := make([]T, r)
	combinationsHelper(input, r, 0, 0, current, &result)

	return result
}

func combinationsHelper[T comparable](input []T, r, index, start int, current []T, result *[][]T) {
	if index == r {
		combination := make([]T, r)
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	for i := start; i < len(input); i++ {
		current[index] = input[i]
		combinationsHelper(input, r, index+1, i+1, current, result)
	}
}
