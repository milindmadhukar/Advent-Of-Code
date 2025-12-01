package utils

import (
	"slices"
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
