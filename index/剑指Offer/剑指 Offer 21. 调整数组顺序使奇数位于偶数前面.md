## 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)

### 解题思路

一道简单应用双指针移动的题，只需要判断左右指针指向的数为奇数还是偶数，然后移动即可。

### 代码

```go
func exchange(nums []int) []int {
	for i, j := 0, len(nums) - 1;i < j; {
        // 碰到左指针是偶数，右指针是奇数，则交换
		if nums[i] % 2 == 0 && nums[j] % 2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i ++
			j --
		}else if nums[i] % 2 == 0 && nums[j] % 2 == 0{ // 如果左指针是偶数，右指针也是偶数，只移动右指针
			j --
		}else if nums[i] % 2 == 1 && nums[j] % 2 == 1{ // 同理
			i ++ 
		}else {
			i ++
			j --
		}
	}
	return nums
}
```

```go
// 优化写法
func exchange(nums []int) []int {
    l, r := 0, len(nums)-1
    for l < r {
        for l < r && nums[l] % 2 == 1 {
            l++
        }
        for l < r && nums[r] % 2 == 0 {
            r--
        }
        nums[l], nums[r] = nums[r], nums[l]
    }
    return nums
}
```