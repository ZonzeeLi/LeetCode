## 剑指 Offer 17. 打印从1到最大的n位数

### 题目传送门

[点击这里](https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

### 解题思路

这道题不考虑大数情况的话就是一道很简单的题，只需要考虑n位数的最大数是多少即可，1位是9，2位是99，3位是999...以此类推，可以得出，n每多1，max要乘10然后加9。这道题正常的做法应该是全排列的做法，应用的是dfs的方法，将每一位进行向下递归统计，只需要考虑首字母不不能是0这种情况即可。

### 代码

```go
func printNumbers(n int) []int {
    var res []int
    var max int
    // 1位是9，2位是99，3位是999，以此类推，可以得出，n每多1，max要乘10然后加9
    for n > 0 {
        max =max*10+9
    	n--
    }
    for i:=1;i<=max;i++{
    	res = append(res,i)
	}
	return res
}

```

```go
func printNumbers(n int) []int {
    res := make([]int, int(math.Pow10(n)-1))
    count := 0
    var dfs func (index int, num []byte, digit int)
    dfs = func (index int, num []byte, digit int) {
                if index == digit {
                    // 如果从'0'开始判断
                    // if num[0] == '0' {
                    //     return 
                    // }
                    tmp, _ := strconv.Atoi(string(num))
                    res[count] = tmp
                    count++
                    return
                }
                for i := '0'; i <= '9'; i++ {
                    num[index] = byte(i)
                    dfs(index + 1, num, digit)
                }
            }
    for digit := 1; digit <= n; digit++ {
        for first := '1'; first <= '9'; first++ {
            num := make([]byte, digit)
            num[0] = byte(first)
            dfs(1 ,num, digit)
        }
    }
    return res
}
```