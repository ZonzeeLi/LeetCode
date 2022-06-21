## 剑指 Offer 51. 数组中的逆序对

### 题目传送门

[点击这里](https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)

### 解题思路

这道题是一道困难题，问题是方法和内部处理的思想不好想。首先这道题用暴力肯定是过不了的。利用的方法就是归并排序，其实归并也就是利用了分治，将一个数组无限的分成两部分..两部分..两部分，最后每一部分是一个数，一个数的话那么不构成逆序对，返回值为0然后，回到上一部分为两个数或者是这种(0,1),(2,2)的情况，由一个数和两个数(两个数需要在分一次，分成两个含有一个数)，(注：这里的数代表的是索引)，然后从两边的第一个数开始比较大小，哪边小哪边的索引就向后移，然后小的数放进新的数组里，(这一步相当于重新排序了)，如果是左边的小的话，就把左边的放进去，然后判断一下右边的当前索引和右边的起始位置差值多少，就是左边当前值造成了多少逆序对，可以这么理解左边是(8,16,22)，右边是(9,18,21),第一次左边的8进去，返回值加0，因为8没有造成逆序，右边的当前索引-起始位置是0，这时左边的索引加1，变成16，比9大， 把9放进去，右边的索引加1，变成18，这时候左边的16小于18，左边16放进去，这时就是左边的数小，然后我们就要判断一下右边的当前索引和右边的起始位置差值多少，这时候差1，就因为右边索引移了，说明之前判断过有比16小的，有多少个小的就是索引和起始位置的差值。当左右两部分有一部分的索引到头了，说明有多余的，需要放进新的数组，如果是右边多余，那就直接放，如果是左边多余，说明还有逆序对，数量就是左边多的个数*右边部分的长度。

### 代码

```go
func reversePairs(nums []int) int {
    // 归并排序
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end - start)/2
    // 分治递归
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid + 1, end)
	tmp := []int{}
    // 双指针
	i, j := start, mid + 1
    // 归并排序
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
            // 这里每一次j移动都会对逆序进行影响，说明nums[i]出现大于nums[j]的情况
            // 在i移动的时候对结果进行影响处理，因为每移动一次i，已经移动的j部分都是逆序的，相当于每次i移动的部分乘以j已经移动的部分，这里也就是1*[j-(mid+1)]
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
        // 同上分析
		cnt += end - mid
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i - start]
	}
	return cnt
}

```