package hashmap

// TwoSum returns the indexes of two different values in nums whose sum is
// target. It returns []int{-1, -1} when no such pair exists.
func TwoSum(nums []int, target int) []int {
	x := make(map[int]int)
	for index := 0; index < len(nums); index++ {
		x[nums[index]] = index
	}

	for now := 0; now < len(nums); now++ {
		index, ok := x[target-nums[now]]
		if ok && index != now {
			return []int{now, index}
		}
	}
	return []int{-1, -1}
}
