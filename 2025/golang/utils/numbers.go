package utils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Gcd[T constraints.Integer](a, b T) T {
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
}

func Lcm[T constraints.Integer](numbers []T) T {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = (result * numbers[i]) / Gcd(result, numbers[i])
	}
	return result
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func Product[T Number](numbers []T) T {
	var product T = 1
	for _, val := range numbers {
		product *= val
	}
	return product
}
