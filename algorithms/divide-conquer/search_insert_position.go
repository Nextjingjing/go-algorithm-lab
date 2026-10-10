package divideconquer

// SearchInsert returns the index of target in nums, or the index where target
// should be inserted to keep nums sorted in ascending order.
//
// nums must be non-empty, sorted in ascending order, and contain distinct
// values. SearchInsert does not modify nums.
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
