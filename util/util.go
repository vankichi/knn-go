package util

// return true if list contains n
func IntContains(list []int32, n int32) bool {
	for _, v := range list {
		if n == v {
			return true
		}
	}
	return false
}

// return true if list contains s
func StrContains(list []string, s string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}
