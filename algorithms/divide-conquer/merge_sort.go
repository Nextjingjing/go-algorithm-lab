package divideconquer

// mergeSort recursively sorts the inclusive range arr[left..right].
//
// d is the slice being sorted.
// left is the first index of the current range.
// right is the last index of the current range.
func mergeSort(
	d []int,
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
// d is the slice being sorted.
// left is the first index of the left half.
// mid is the last index of the left half.
// right is the last index of the right half.
func merge(
	d []int,
	left int,
	mid int,
	right int,
) {
	// merged temporarily stores the sorted values before copying them back.
	merged := []int{}

	// i walks through the left half.
	i := left

	// j walks through the right half.
	j := mid + 1

	for i <= mid && j <= right {
		if d[i] >= d[j] {
			merged = append(merged, d[j])
			j++
		} else {
			merged = append(merged, d[i])
			i++
		}
	}

	for ; i <= mid; i++ {
		merged = append(merged, d[i])
	}

	for ; j <= right; j++ {
		merged = append(merged, d[j])
	}

	// k is the index inside merged. left+k maps it back to the original slice.
	for k := 0; k < len(merged); k++ {
		d[left+k] = merged[k]
	}
}

// MergeSort sorts d in place.
//
// d is the slice that should be sorted.
func MergeSort(d []int) {
	// left and right describe the full inclusive range of the slice.
	left := 0
	right := len(d) - 1
	mergeSort(d, left, right)
}
