package main

// 238. 除自身以外数组的乘积
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
func productExceptSelf(nums []int) []int {
	// leftRes[i]和rightRes[i]分别表示，nums[i]左边的数乘积和右边的数乘积，即最终结果为leftRes[i] * rightRes[i]
	leftRes := make([]int, len(nums))
	rightRes := make([]int, len(nums))

	leftRes[0] = 1
	rightRes[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		leftRes[i] = leftRes[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		rightRes[i] = rightRes[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = leftRes[i] * rightRes[i]
	}

	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	for _, v := range result {
		println(v) // 输出: 24 12 8 6
	}
}
