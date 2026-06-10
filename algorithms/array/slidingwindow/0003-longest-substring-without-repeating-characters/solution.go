package solution

func lengthOfLongestSubstring(s string) int {
	left := 0
	res := 0
	count := make([]int, 128)

	for right := 0; right < len(s); right++ {
		count[s[right]]++

		for count[s[right]] > 1 {
			count[s[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}
