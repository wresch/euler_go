package euler_go


func IntFreq(slice []int64) (map[int64]int64) {
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