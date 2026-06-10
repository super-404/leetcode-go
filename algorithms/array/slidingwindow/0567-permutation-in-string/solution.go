package solution

func checkInclusion(s1 string, s2 string) bool {
	left := 0
	right := 0
	window := make([]int, 1024)
	tArr := make([]int, 1024)
	for _, v := range s1 {
		tArr[v]++
	}
	for right < len(s2) {
		window[s2[right]]++
		if right-left+1 == len(s1) {
			if isSame(window, tArr) {
				return true
			} else {
				window[s2[left]]--
				left++
			}
		}
		right++
	}
	return false
}
func isSame(window []int, target []int) bool {
	for i, v := range target {
		if window[i] != v {
			return false
		}
	}
	return true
}
