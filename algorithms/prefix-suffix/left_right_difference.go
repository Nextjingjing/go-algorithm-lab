package prefixsuffix

// LeftRightDifference returns, for each index, the absolute difference
// between the sum of values to its left and the sum of values to its right.
// It returns a result with the same length as nums and does not modify nums.
func LeftRightDifference(nums []int) []int {
	answer := []int{}

	leftSum := []int{0}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		leftSum = append(leftSum, sum)
	}

	total := leftSum[len(nums)]
	for i := 0; i < len(nums); i++ {
		right := total - nums[i] - leftSum[i]
		diff := leftSum[i] - right
		if diff < 0 {
			diff *= -1
		}
		answer = append(answer, diff)
	}
	return answer
}
