package polyfill

func Map[T any, R any](input []T, f func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Find[T comparable](input []T, item T) int {
	index := -1
	for i, v := range input {
		if v == item {
			index = i
			break
		}
	}
	return index
}
