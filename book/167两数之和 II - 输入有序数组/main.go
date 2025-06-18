package main

// 167. 两数之和 II - 输入有序数组
//给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
//
//以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
//
//你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
//
//你所设计的解决方案必须只使用常量级的额外空间。

//func twoSum(numbers []int, target int) []int {
//	for i := 0; i < len(numbers); i++ {
//		for j := i + 1; j < len(numbers); j++ {
//			if numbers[i]+numbers[j] == target {
//				return []int{i + 1, j + 1}
//			} else if numbers[i]+numbers[j] > target {
//				break
//			}
//		}
//	}
//	return []int{0, 0}
//}

// 更好的做法
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	var exists = make(map[int]int, n)
	for i := 0; i < n; i++ {
		exists[numbers[i]] = i + 1
	}
	for i := 0; i < n; i++ {
		d := target - numbers[i]
		if j, ok := exists[d]; ok {
			return []int{i + 1, j}
		}
	}
	return nil
}
