// LeetCode 66. Plus One 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
package solution

// plusOne 在 digits 表示的数上加一
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 全为 9，需要进位到新的一位
	return append([]int{1}, digits...)
}
