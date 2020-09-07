package utils

func InSlice(a []string, x string) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return false
}

func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
