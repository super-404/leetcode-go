package solution

func findMaxAverage(nums []int, k int) float64 {
	if k <= 0 || k > len(nums) {
		return 0
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum
	for right := k; right < len(nums); right++ {
		windowSum += nums[right] - nums[right-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return float64(maxSum) / float64(k)
}
