package bruteforce

// MinValue returns the smallest value in d.
//
// d must contain at least one value.
func MinValue(d []int) int {
	if len(d) == 0 {
		panic("MinValue: empty slice")
	}

	_min := d[0]
	for i := 1; i < len(d); i++ {
		_min = min(_min, d[i])
	}

	return _min
}
