## 剑指 Offer 38. 字符串的排列

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)

### 解题思路

首先分析这道题，这道题要把字符串的所有排列都统计出来，但是有一个问题会出现重复的元素，这样子如果单纯的遍历的话会出现同样的结果，所以这里要做记忆化。这道题可以使用回溯的算法，回溯的算法比较麻烦，难度倒是不难，麻烦在需要两个记忆化，一个记录是否添加过当前索引的字符，一个是要记录该字符串是否添加过结果里，不过这也是一个可行的方法。对于上述我们可以进一步优化，只用一个记忆化，多加一条规则即可，即对字符串排序，将重复的子串放在一起，要添加该字符，只添加重复子串的第一个。另外，这道题还有一种解法，即下一个排列，我们可以将字符串按照规定的字典序进行排列，即得到最小的字典序，然后一直寻找下一个排列即可，至于[31. 下一个排列这道题](https://leetcode-cn.com/problems/next-permutation/) 可以看一下我写的[31. 下一个排列的题解](https://github.com/ZonzeeLi/LeetCode/blob/master/index/31-40/31.%20%E4%B8%8B%E4%B8%80%E4%B8%AA%E6%8E%92%E5%88%97.md)。

### 代码

```go
func permutation(s string) []string {
    b := []byte(s)
    var res []string
    n := len(b)
    memo := make([]bool, n)
    memostr := make(map[string]bool)
    // dfs + 回溯
    var dfs func(tmp []byte, memo []bool) 
    dfs = func(tmp []byte, memo []bool) {
        if len(tmp) == n {
            if memostr[string(tmp)] {
                return
            }
            t := make([]byte, len(tmp))
		    copy(t, tmp)
            res = append(res, string(t))
            memostr[string(t)] = true
            return
        }
        for i := 0;i < n;i ++ {
            if memo[i] {
                continue
            }
            memo[i] = true
            tmp = append(tmp, b[i])
            dfs(tmp, memo)
            tmp = tmp[:len(tmp)-1]
            memo[i] = false
        }
    }
    dfs([]byte{}, memo)
    return res
}
```

```go
// 回溯优化
func permutation(s string) (ans []string) {
    t := []byte(s)
    // 排序，主要目的是为了将重复的字符放在一起
    sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
    n := len(t)
    perm := make([]byte, 0, n)
    vis := make([]bool, n)
    var backtrack func(int)
    backtrack = func(i int) {
        if i == n {
            ans = append(ans, string(perm))
            return
        }
        for j, b := range vis {
            // 限制条件，当前字符一定是重复子串的第一个，比如“abb”，我们只加入第一个“b”，第二个“b”在第一个“b”没有加入时不要加入。
            if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
                continue
            }
            vis[j] = true
            perm = append(perm, t[j])
            backtrack(i + 1)
            perm = perm[:len(perm)-1]
            vis[j] = false
        }
    }
    backtrack(0)
    return
}
```

```go
func reverse(a []byte) {
    for i, n := 0, len(a); i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}

func nextPermutation(nums []byte) bool {
    n := len(nums)
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    if i < 0 {
        return false
    }
    j := n - 1
    for j >= 0 && nums[i] >= nums[j] {
        j--
    }
    nums[i], nums[j] = nums[j], nums[i]
    reverse(nums[i+1:])
    return true
}

func permutation(s string) (ans []string) {
    t := []byte(s)
    sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
    for {
        ans = append(ans, string(t))
        if !nextPermutation(t) {
            break
        }
    }
    return
}
```
