package utils

func Permutations[T comparable](input []T, size int) <-chan []T {
	ch := make(chan []T)
	go func() {
		defer close(ch)
		current := make([]T, size)
		permutationsHelper(input, size, 0, current, ch)
	}()
	return ch
}

func permutationsHelper[T comparable](input []T, size int, position int, current []T, ch chan<- []T) {
	if position == size {
		combination := make([]T, size)
		copy(combination, current)
		ch <- combination
		return
	}

	for i := 0; i < len(input); i++ {
		current[position] = input[i]
		permutationsHelper(input, size, position+1, current, ch)
	}
}

func Combinations[T comparable](input []T, r int) <-chan []T {
	ch := make(chan []T)
	go func() {
		defer close(ch)
		current := make([]T, r)
		combinationsHelper(input, r, 0, 0, current, ch)
	}()
	return ch
}

func combinationsHelper[T comparable](input []T, r, index, start int, current []T, ch chan<- []T) {
	if index == r {
		combination := make([]T, r)
		copy(combination, current)
		ch <- combination
		return
	}

	for i := start; i < len(input); i++ {
		current[index] = input[i]
		combinationsHelper(input, r, index+1, i+1, current, ch)
	}
}
