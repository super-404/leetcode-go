package solution

func longestOnes(nums []int, k int) int {
	left := 0
	right := 0
	res := 0
	count := 0
	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}
		for count > k {
			if nums[left] == 0 {
				count--
			}
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}
