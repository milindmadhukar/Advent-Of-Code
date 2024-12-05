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

func IndexOf[K comparable](slice []K, element K) int {
	for i, item := range slice {
		if item == element {
			return i
		}
	}
	return -1
}

func Contains[K comparable](slice []K, element K) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
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

func CountOfAll[K comparable](slice []K) map[K]int {
	counts := make(map[K]int)
	for _, item := range slice {
		counts[item]++
	}
	return counts
}

func Remove[K comparable](slice []K, element K) []K {
	var result []K
	for _, item := range slice {
		if item != element {
			result = append(result, item)
		}
	}
	return result
}

func Pop[K any](slice []K, index int) (K, []K) {
	if index < 0 || index >= len(slice) {
		panic("Index out of range")
	}

	result := slice[index]
	slice = slices.Delete(slice, index, index+1)
	return result, slice
}

func Insert[K any](slice []K, index int, element K) []K {
	if index < 0 || index > len(slice) {
		panic("Index out of range")
	}

	slice = append(slice, element)
	copy(slice[index+1:], slice[index:])
	slice[index] = element
	return slice
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

func Reduce[K any, V any](reducer func(V, K) V, values []K, initialValue V) V {
	result := initialValue

	for _, value := range values {
		result = reducer(result, value)
	}

	return result
}
