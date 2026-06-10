package solution

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 1
	j := 0
	for j < len(nums) && i < len(nums) {
		if nums[j] != nums[i] {
			nums[j+1] = nums[i]
			j++
		}
		i++
	}
	return j
}
