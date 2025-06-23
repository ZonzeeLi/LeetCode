package main

import "sort"

// 15. 三数之和
//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。

// 这代码写的太丑了
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	k := 0
	for k < len(nums)-2 {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			k++
			continue
		}
		i := k + 1
		j := len(nums) - 1
		for j > i {
			if nums[k]+nums[i]+nums[j] == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[k]+nums[i]+nums[j] < 0 {
				i++
			} else {
				j--
			}
		}
		k++
	}
	return res
}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//result := threeSum(nums)
	//for _, triplet := range result {
	//	println(triplet[0], triplet[1], triplet[2])
	//}

	nums := []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}
	result := threeSum(nums)
	for _, triplet := range result {
		println(triplet[0], triplet[1], triplet[2])
	}
}
