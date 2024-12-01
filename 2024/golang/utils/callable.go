package utils

func MustClosure[K any, L any](callable func(L) (K, error)) func(L) K {
	return func(l L) K {
		output, err := callable(l)
		if err != nil {
			panic(err)
		}

		return output
	}
}

func Map[K any, V any](mapper func(K) V, values []K) []V {
	var result []V

	for _, value := range values {
		result = append(result, mapper(value))
	}

	return result
}

func Filter[K any](predicate func(K) bool, values []K) []K {
  var result []K

  for _, value := range values {
    if predicate(value) {
      result = append(result, value)
    }
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

func Contains[K comparable](slice []K, element K) bool {
  for _, item := range slice {
    if item == element {
      return true
    }
  }
  return false
}
