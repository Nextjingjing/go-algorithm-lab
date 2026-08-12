package twopointers

// MoveZeroes moves all zero values in nums to the end in place.
//
// The relative order of non-zero values must remain unchanged.
func MoveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
