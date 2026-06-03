// LeetCode 1536. Minimum Swaps to Arrange a Binary Grid 排布二进制网格的最少交换次数
// 给定 n×n 二进制网格，每次可交换相邻两行。若主对角线上方的格子均为 0 则有效。
// 返回使网格有效所需的最少交换次数，若不可能则返回 -1。
package solution

// minSwaps 返回使 grid 有效所需的最少交换次数
// 思路：在每一行的右上角若发现 1，则在该列中找「能合法放在当前行」的最近的 0（即该行最右 1 的列 ≤ 当前行号），逐行交换并记录次数，重复直到该行右上角全为 0。
func minSwaps(grid [][]int) int {
	n := len(grid)
	// pos[i]：第 i 行最右 1 的列下标，用于判断该行在「哪一列」有 1，以及能否放在某行
	pos := make([]int, n)
	for i := range pos {
		pos[i] = -1
	}
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				pos[i] = j
				break
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		// 若当前行右上角有 1（pos[i] > i），则要在「这一列」找到最近的 0 所在的行并换上来
		for pos[i] > i {
			// 列号就是 pos[i]；要找该列中从第 i 行起最近的、且能合法放在第 i 行的 0
			// 「该列是 0」等价于该行最右 1 的列 < pos[i]，即 pos[k] <= pos[i]-1
			// 「能合法放在第 i 行」等价于该行最右 1 的列 <= i，即 pos[k] <= i
			// 因此需要 pos[k] <= i（满足则自动在列 pos[i] 为 0）
			k := -1
			for j := i; j < n; j++ {
				if pos[j] <= i {
					k = j
					break
				}
			}
			if k == -1 {
				return -1
			}
			// 把第 k 行逐行交换到第 i 行
			for t := k; t > i; t-- {
				pos[t], pos[t-1] = pos[t-1], pos[t]
				ans++
			}
			// 交换后第 i 行已是「带 0 的那一行」，其最右 1 的列 <= i，故不再有右上角 1，可继续下一行
		}
	}
	return ans
}
