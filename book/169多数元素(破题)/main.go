package main

// 169. 多数元素
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
func majorityElement(nums []int) int {
	var count, curnum int
	for i := 0; i < len(nums); i++ {
		if curnum == 0 || count == 0 {
			curnum = nums[i]
			count++
		} else if curnum == nums[i] {
			count++
		} else if curnum != nums[i] {
			count--
		}
	}
	return curnum
}
