package slice

// Contains returns positive boolean if slice contains string
func Contains(s string, ss []string) bool {
	for _, n := range ss {
		if s == n {
			return true
		}
	}
	return false
}
