package arrays

func Sum(arr []int) int {
	sum := 0

	for _, value := range arr {
		sum += value
	}
	return sum

}
func MultipleArray(numb ...[]int) []int {
	length := len(numb)
	sums := make([]int, length)

	for i, numbers := range numb {
		sums[i] = Sum(numbers)
	}

	return sums
}
