package prefixsuffix

// FirstTryProductExceptSelf returns a slice where result[i] is the product of every
// value in nums except nums[i].
//
// nums must contain at least two values. The input is not modified.
func FirstTryProductExceptSelf(nums []int) []int {
	rightProducts := make(map[int]int)
	product := 1
	for i := len(nums) - 1; i > 0; i-- {
		product *= nums[i]
		rightProducts[i-1] = product
	}
	r := []int{}
	r = append(r, rightProducts[0])
	leftProduct := 1
	for i := 1; i < len(nums); i++ {
		leftProduct *= nums[i-1]
		if i < len(nums)-1 {
			r = append(r, leftProduct*rightProducts[i])
		}
	}
	r = append(r, leftProduct)
	return r
}

// ProductExceptSelf returns a slice where result[i] is the product of every
// value in nums except nums[i].
//
// nums must contain at least two values. The input is not modified.
func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	leftProduct := 1
	for i := 0; i < len(nums); i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}
	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return result
}
