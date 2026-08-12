package divideconquer

// binarySearch searches for x inside the inclusive range d[left..right].
//
// d is the sorted slice being searched.
// x is the value to find.
// left is the first index of the current search range.
// right is the last index of the current search range.
func binarySearch(
	d []int,
	x int,
	left int,
	right int,
) int {
	// mid is the index in the middle of the current search range.
	mid := (right + left) / 2

	if d[mid] == x {
		return mid
	}

	if mid == 0 {
		return -1
	}

	if d[mid] > x {
		return binarySearch(d, x, left, mid-1)
	} else {
		return binarySearch(d, x, mid+1, right)
	}
}

// BinarySearch searches for x in d and returns its index.
//
// d must be sorted in ascending order.
// x is the value to find.
// It returns -1 when x is not found.
func BinarySearch(
	d []int,
	x int,
) int {
	// left and right describe the full inclusive range of the slice.
	left := 0
	right := len(d) - 1
	return binarySearch(d, x, left, right)
}
