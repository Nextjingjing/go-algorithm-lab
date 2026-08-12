package divideconquer

// mergeSort recursively sorts the inclusive range arr[left..right].
//
// d points to the slice being sorted.
// left is the first index of the current range.
// right is the last index of the current range.
func mergeSort(
	d *[]int,
	left int,
	right int,
) {
	if left >= right {
		return
	}

	// mid splits the current range into left and right halves.
	mid := (left + right) / 2
	mergeSort(d, left, mid)
	mergeSort(d, mid+1, right)
	merge(d, left, mid, right)
}

// merge combines two already-sorted inclusive ranges:
//
// left half:  arr[left..mid]
// right half: arr[mid+1..right]
//
// d points to the slice being sorted.
// left is the first index of the left half.
// mid is the last index of the left half.
// right is the last index of the right half.
func merge(
	d *[]int,
	left int,
	mid int,
	right int,
) []int {
	// arr is the actual slice value behind the pointer d.
	arr := *d

	// merged temporarily stores the sorted values before copying them back.
	merged := []int{}

	// i walks through the left half.
	i := left

	// j walks through the right half.
	j := mid + 1

	for i <= mid && j <= right {
		if arr[i] >= arr[j] {
			merged = append(merged, arr[j])
			j++
		} else {
			merged = append(merged, arr[i])
			i++
		}
	}

	for ; i <= mid; i++ {
		merged = append(merged, arr[i])
	}

	for ; j <= right; j++ {
		merged = append(merged, arr[j])
	}

	// k is the index inside merged. left+k maps it back to the original slice.
	for k := 0; k < len(merged); k++ {
		arr[left+k] = merged[k]
	}
	return arr
}

// MergeSort sorts d in place.
//
// d points to the slice that should be sorted.
func MergeSort(d *[]int) {
	// left and right describe the full inclusive range of the slice.
	left := 0
	right := len(*d) - 1
	mergeSort(d, left, right)
}
