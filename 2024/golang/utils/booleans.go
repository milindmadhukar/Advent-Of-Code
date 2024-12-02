package utils

func All[T any](arr []T, predicate func(T) bool) bool {
	for _, val := range arr {
		if !predicate(val) {
			return false
		}
	}
	return true
}

func Any[T any](arr []T, predicate func(T) bool) bool {
  for _, val := range arr {
    if predicate(val) {
      return true
    }
  }
  return false
}
