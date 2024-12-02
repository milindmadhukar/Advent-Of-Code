package utils

import (
	"slices"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/models"
)

func GetUniqueElements[K comparable](slice []K) []K {
	var uniqueElements []K
	for _, element := range slice {
		if !slices.Contains(uniqueElements, element) {
			uniqueElements = append(uniqueElements, element)
		}
	}
	return uniqueElements
}

func Intersection[K comparable](A, B []K) []K {
	var result []K
	for _, item := range B {
		if slices.Contains(A, item) {
			result = append(result, item)
		}
	}

	return result
}

func CountOf[K comparable](slice []K, element K) int {
	count := 0
	for _, item := range slice {
		if item == element {
			count++
		}
	}
	return count
}

func Zip[T, U any](slice1 []T, slice2 []U) []models.Pair[T, U] {
	minLen := len(slice1)
	if len(slice2) < minLen {
		minLen = len(slice2)
	}

	result := make([]models.Pair[T, U], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = models.Pair[T, U]{First: slice1[i], Second: slice2[i]}
	}

	return result
}

func ZipLongest[T, U any](slice1 []T, slice2 []U) []models.Pair[T, U] {
	maxLen := len(slice1)
	if len(slice2) > maxLen {
		maxLen = len(slice2)
	}

	result := make([]models.Pair[T, U], maxLen)
	for i := 0; i < maxLen; i++ {
		var first T
		if i < len(slice1) {
			first = slice1[i]
		}

		var second U
		if i < len(slice2) {
			second = slice2[i]
		}

		result[i] = models.Pair[T, U]{First: first, Second: second}
	}

	return result
}
