## 剑指 Offer 03. 数组中重复的数字

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)

### 解题思路

一道简单题，方法有很多。这道题的大体思路有排序后，比对前后值，找到重复数字，还可以使用hashmap用来存放数字的出现次数，另外有一种非常节省空间的算法，原地置换。解释一下该算法的思路，该算法是利用了条件所有数字在0~n-1范围内，所以可以保证如果不出现重复数字的话，那么就会`nums[i]=i`，所以一但不等于的话，我们就可以让`nums[nums[i]]`与`nums[i]`进行置换，这样子能够保证`nums[nums[i]] = nums[i]`，然后继续置换直到在置换前就已经满足`nums[nums[i]] = nums[i]`。

### 代码

```go
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[nums[i]]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
```

```go
func findRepeatNumber(nums []int) int {
	// 该算法利用哈希表 时间复杂度为O(n) 空间复杂度为O(n)
	c := make(map[int]bool, len(nums))
	var b int
	for i := 0; i < len(nums); i++ {
		if c[nums[i]] {
			b = nums[i]
			break
		}
		c[nums[i]] = true
	}
	return b
}
```