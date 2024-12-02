package utils

import "golang.org/x/exp/constraints"

type Number interface {
    constraints.Integer | constraints.Float
}


func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
}

func Lcm(numbers []int) int {
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


