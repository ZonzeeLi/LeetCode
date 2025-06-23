package main

// 209. 长度最小的子数组
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

func minSubArrayLen(target int, nums []int) int {
	res := 999999
	// 先统计一个前缀和
	headSum := make([]int, len(nums))
	headSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		headSum[i] = headSum[i-1] + nums[i]
	}
	if headSum[len(nums)-1] < target {
		return 0
	}
	// 从第一个数开始，每次比较和target的大小是否一致，一样则直接返回
	// 如果当前在索引j，如果大，因为是正数，所以可以移动左端点，即分数是headSum[j] - headSum[i]
	// 如果小，则移动右端点
	j := 0
	for i := 0; i < len(nums); i++ {
		for j < len(nums) {
			if headSum[j] >= target {
				res = min(res, j+1)
			}
			if headSum[j]-headSum[i] >= target {
				res = min(res, j-i)
				break
			} else if headSum[j]-headSum[i] < target {
				j++
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	// Example usage
	//target := 7
	//nums := []int{2, 3, 1, 2, 4, 3}
	//result := minSubArrayLen(target, nums)
	//println(result) // Output: 2 (subarray [4,3] has sum >= 7)

	target := 15
	nums := []int{1, 2, 3, 4, 5}
	result := minSubArrayLen(target, nums)
	println(result)
}
