package rnm

func isSameSlices(s1 []string, s2 []string) bool {
	if !(s1 != nil && s2 != nil && len(s1) == len(s2)) {
		return false
	}

	for i, v1 := range s1 {
		if v1 != s2[i] {
			return false
		}
	}

	return true
}
