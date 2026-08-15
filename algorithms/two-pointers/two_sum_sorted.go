package twopointers

// TwoSumSorted returns the 1-based indexes of two values in numbers whose sum
// equals target.
//
// numbers must be sorted in ascending order. The input is not modified, and
// the problem contract guarantees exactly one matching pair.
func TwoSumSorted(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
