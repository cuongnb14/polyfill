package polyfill

func Ternary[T any](condition bool, trueResult, falseResult T) T {
	if condition {
		return trueResult
	}
	return falseResult
}
