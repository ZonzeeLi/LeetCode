package main

// 189. 轮转数组
//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//进阶：
//
//尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
//你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

// 很笨的解法
//func rotate(nums []int, k int) {
//	if len(nums) <= 1 {
//		return
//	}
//	k = k % len(nums)
//	newNums := make([]int, len(nums))
//	newNums = append([]int{}, nums[len(nums)-k:len(nums)]...)
//	newNums = append(newNums, nums[:len(nums)-k]...)
//	copy(nums, newNums)
//}

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	for _, v := range nums {
		println(v)
	}
}
