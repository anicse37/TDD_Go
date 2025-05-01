package arrays

func Reduce[A any](Collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range Collection {
		result = f(result, x)
	}
	return result
}
