package prefixsuffix

// LargestAltitude returns the highest altitude reached when starting at
// altitude 0 and applying each gain in order. The input is not modified.
// An empty gain slice returns 0.
func LargestAltitude(gain []int) int {
	d := []int{0}
	now := 0
	largestAltitude := 0
	for i := 0; i < len(gain); i++ {
		d = append(d, now+gain[i])
		now = d[len(d)-1]
		largestAltitude = max(now, largestAltitude)
	}
	return largestAltitude
}
