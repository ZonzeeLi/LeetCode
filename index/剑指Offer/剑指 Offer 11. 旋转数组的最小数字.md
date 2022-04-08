## 剑指 Offer 11. 旋转数组的最小数字

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

### 解题思路

这道题是用二分法来做，但其实二分法要求的是顺序性，这道题存在顺序性吗？其实是存在的，这道题可以看成是两段有序的结合在了一起。情况大概只有这三种，举例如下：

1. `[4 5 0 1 2 3]`，这种情况是`mid < end`，这说明`mid`之后的肯定有序的，所以最小值只可能是包含此时`mid`的前半部分，所以我们让`end = mid`
2. `[3 4 5 0 1 2]`，这种情况是`mid > end`，这说明`mid`之前的肯定是有序的，也肯定是反转过的，不然前面的不会比后面的大，所以最小值在不包含`mid`的之后不分，所以让`start = mid+1`
3. `[2 2 2 0 1 2]`和`[2 0 1 2 2 2 2]`，这种情况是`mid = end`，这种情况其实是没办法判断小值在哪一边，我们只能缩短长度，让他尽可能的变成1和2这两种情况。

### 代码

```go
func minArray(numbers []int) int {
    start, mid, end := 0, 0, len(numbers) - 1
    for start < end {
        mid = start+(end-start)/2
        if numbers[mid] > numbers[end] {
            start = mid + 1
        }else if numbers[mid] < numbers[end] {
            end = mid
        }else if numbers[mid] == numbers[end] {
            end --
        }
    }
    return numbers[start]
}
```