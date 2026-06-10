# LeetCode-Go

使用 Go 记录 LeetCode 刷题代码、测试用例和常用数据结构。

## 目录结构

```text
LeetCode-Go/
├── README.md
├── go.mod
├── datastructures/         # 常用数据结构定义
│   ├── linkedlist.go
│   └── tree.go
└── algorithms/             # 按题型分类存放题解
    ├── array/
    │   ├── 0066-plus-one/
    │   ├── slidingwindow/
    │   │   ├── 0003-longest-substring-without-repeating-characters/
    │   │   ├── 0076-minimum-window-substring/
    │   │   ├── 0567-permutation-in-string/
    │   │   ├── 0643-maximum-average-subarray-i/
    │   │   ├── 0674-longest-continuous-increasing-subsequence/
    │   │   ├── 1004-max-consecutive-ones-iii/
    │   │   └── 1343-number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
    │   └── twopointers/
    │       ├── 0005-longest-palindromic-substring/
    │       ├── 0026-remove-duplicates-from-sorted-array/
    │       ├── 0027-remove-element/
    │       ├── 0167-two-sum-ii-input-array-is-sorted/
    │       ├── 0283-move-zeroes/
    │       └── 0344-reverse-string/
    ├── linkedlist/
    │   ├── 0019-remove-nth-node-from-end-of-list/
    │   ├── 0021-merge-two-sorted-lists/
    │   ├── 0023-merge-k-sorted-lists/
    │   ├── 0141-linked-list-cycle/
    │   ├── 0142-linked-list-cycle-ii/
    │   ├── 0160-intersection-of-two-linked-lists/
    │   ├── 0086-partition-list/
    │   └── 0876-middle-of-the-linked-list/
    ├── greedy/
    │   └── 1536-minimum-swaps-to-arrange-a-binary-grid/
    ├── binarytree/
    ├── backtrace/
    ├── dp/
    └── sorts/
```

## 命名规则

- 题目目录使用 `题号4位-kebab-case`，例如 `0023-merge-k-sorted-lists`
- 每道题目录内使用 `solution.go` 和可选的 `solution_test.go`
- 题目包名统一使用 `package solution`
- 公共链表、树等结构放在 `datastructures`

## 运行测试

```bash
go test ./...
```

## 已整理题目

| 题号 | 题目 | 分类 |
| --- | --- | --- |
| 3 | Longest Substring Without Repeating Characters | array/slidingwindow |
| 5 | Longest Palindromic Substring | array/twopointers |
| 19 | Remove Nth Node From End of List | linkedlist |
| 21 | Merge Two Sorted Lists | linkedlist |
| 23 | Merge k Sorted Lists | linkedlist |
| 26 | Remove Duplicates from Sorted Array | array/twopointers |
| 27 | Remove Element | array/twopointers |
| 66 | Plus One | array |
| 76 | Minimum Window Substring | array/slidingwindow |
| 86 | Partition List | linkedlist |
| 141 | Linked List Cycle | linkedlist |
| 142 | Linked List Cycle II | linkedlist |
| 160 | Intersection of Two Linked Lists | linkedlist |
| 167 | Two Sum II - Input Array Is Sorted | array/twopointers |
| 283 | Move Zeroes | array/twopointers |
| 344 | Reverse String | array/twopointers |
| 567 | Permutation in String | array/slidingwindow |
| 643 | Maximum Average Subarray I | array/slidingwindow |
| 674 | Longest Continuous Increasing Subsequence | array/slidingwindow |
| 876 | Middle of the Linked List | linkedlist |
| 1004 | Max Consecutive Ones III | array/slidingwindow |
| 1343 | Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold | array/slidingwindow |
| 1536 | Minimum Swaps to Arrange a Binary Grid | greedy |
