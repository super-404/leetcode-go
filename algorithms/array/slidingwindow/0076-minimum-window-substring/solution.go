package solution

func minWindow(s string, t string) string {
	left := 0
	right := 0
	window := make([]int, 1024)
	res := 100000001
	tArr := make([]int, 1024)
	start := 0
	end := 0
	for _, v := range t {
		tArr[v]++
	}
	for right < len(s) {
		window[s[right]]++
		for isSame(window, tArr) {
			if res > right-left+1 {
				res = right - left + 1
				start = left
				end = right + 1
			}
			window[s[left]]--
			left++
		}
		right++
	}
	return s[start:end]
}
func isSame(window []int, target []int) bool {
	for i, v := range target {
		if window[i] < v {
			return false
		}
	}
	return true
}
