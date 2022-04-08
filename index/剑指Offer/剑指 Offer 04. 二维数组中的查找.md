## 剑指 Offer 04. 二维数组中的查找

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)

### 解题思路

这道题的想法不难，首先可以使用暴力解法，但这当然不是最好的方法。可以转换一下思想，用线性查找或者是二分查找，线性查找是利用了每一行和每一列都递增的规律，而且既然是有序的，那么同样也可以使用二分查找来做。

### 代码

```go
func findNumberIn2DArray(matrix [][]int, target int) bool {
    // 重点是要找到起始行
    line := len(matrix) - 1
    row := 0
    for line >= 0 && row < len(matrix[0]) {
        if matrix[line][row] == target {
            return true
        }else if matrix[line][row] > target {
            // 如果大于说明向右的都大于，要向上一行查找
            line --
        }else if matrix[line][row] < target {
            row ++
        }
    }
    return false
}
```

```go
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		 l :=len(matrix[i])
		if l>0 &&matrix[i][0] <= target && target <= matrix[i][l-1] {
			ok := BinarySearch(matrix[i], target)
			if ok {
				return true
			}
		}
	}

	return false
}

func BinarySearch(a []int, target int) bool {
	l := 0
	r := len(a) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if a[mid] == target {
			return true
		} else if a[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
```