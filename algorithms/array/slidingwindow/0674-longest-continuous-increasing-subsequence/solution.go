package solution

func findLengthOfLCIS(nums []int) int {
	right := 0
	res := 0
	l := 0
	for right < len(nums) {
		l++
		if right > 0 && nums[right-1] >= nums[right] {
			l = 1
		}
		res = max(res, l)
		right++
	}
	return res
}
