package check_inclusion

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1map := make(map[string]int)
	s2map := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		s1map[string(s1[i])]++
		s2map[string(s2[i])]++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if isSame(s1map, s2map) {
			return true
		}
		s2map[string(s2[i])]--
		s2map[string(s2[i+len(s1)])]++
	}
	return isSame(s1map, s2map)
}

func isSame(s1map, s2map map[string]int) bool {
	for k, v := range s1map {
		if s2map[k] != v {
			return false
		}
	}
	return true
}
