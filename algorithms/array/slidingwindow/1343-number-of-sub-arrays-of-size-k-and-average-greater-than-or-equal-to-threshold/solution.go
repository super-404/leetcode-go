package solution

func numOfSubarrays(arr []int, k int, threshold int) int {
	left := 0
	right := 0
	// windows := make([]int,k)
	sum := 0
	res := 0
	for right < len(arr) {
		sum += arr[right]
		right++
		if right-left == k && sum/k >= threshold {
			res++
			sum -= arr[left]
		}
	}
	return res
}
