package main

import "fmt"

// 88. 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
// 题目描述：给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

func merge(nums1 []int, m int, nums2 []int, n int) {
	arr := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n && k < m+n {
		if nums1[i] <= nums2[j] {
			arr[k] = nums1[i]
			i++
		} else {
			arr[k] = nums2[j]
			j++
		}
		k++
	}
	if i < m {
		for i < m {
			arr[k] = nums1[i]
			i++
			k++
		}
	}
	if j < n {
		for j < n {
			arr[k] = nums2[j]
			j++
			k++
		}
	}
	copy(nums1, arr)
}

// 双指针法，从后往前合并
func merge2(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if index1 < 0 {
			nums1[i] = nums2[index2]
			index2--
		} else if index2 < 0 {
			nums1[i] = nums1[index1]
			index1--
		} else if nums1[index1] > nums2[index2] {
			nums1[i] = nums1[index1]
			index1--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
