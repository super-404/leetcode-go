package solution

func moveZeroes(nums []int) {
	i := 0
	j := 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
		i++
	}
}
