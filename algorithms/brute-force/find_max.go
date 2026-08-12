package bruteforce

// FindMax returns the largest value in d.
//
// d points to the slice being searched.
// The slice must contain at least one value.
func FindMax(d *[]int) int {
	// arr is the actual slice value behind the pointer d.
	arr := *d

	// m stores the largest value found so far.
	m := arr[0]

	// i walks through the remaining values after the first one.
	for i := 1; i < len(arr); i++ {
		m = max(m, arr[i])
	}

	return m
}
