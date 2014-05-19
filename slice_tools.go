package euler_go

func IntFreq(slice []int64) map[int64]int64 {
	m := make(map[int64]int64)
	for _, i := range slice {
		if n, ok := m[i]; ok {
			m[i] = n + 1
		} else {
			m[i] = 1
		}
	}
	return m
}

func SliceProd(slice []int64) int64 {
	var r int64 = 1
	for _, i := range slice {
		r *= i
	}
	return r
}
