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
