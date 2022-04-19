## 剑指 Offer 13. 机器人的运动范围

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

### 解题思路

这道题又是那种岛屿类型问题，经典的bfs模板题，没什么大难度，只是增加了一条额外限制条件，其他的完全一样。

### 代码

```go
var dir = [][]int{
    // 看了官方解的优化，这里不需要计算向上和向左的，因为从[0,0]开始bfs，到达某一点都是上一个点向右向下遍历过来的，不需要再计算向上向左的计算
    {1, 0},
    {0, 1},
}

func movingCount(m int, n int, k int) int {
    // 经典的bfs模板题
    // 这道题只是增加了一个不能继续bfs的限制条件，即x和y的数位和要不能大于k
    if k == 0 {
        return 1
    }

    vis := make([][]bool, m)
    for i := range vis {
        vis[i] = make([]bool, n)
    }
    queue := [][]int{}
    queue = append(queue, []int{0, 0})
    vis[0][0] = true
    res := 1
    for len(queue) > 0 {
        q := queue[0]
        queue = queue[1:]
        for _, v := range dir {
            nx := q[0] + v[0]
            ny := q[1] + v[1]
            if nx < 0 || nx >= m || ny < 0 || ny >= n || vis[nx][ny] || get(nx, ny) > k {
                continue
            }
            queue = append(queue, []int{nx, ny})
            vis[nx][ny] = true
            res ++
        }
    }
    return res
}

func get(x, y int) int {
    var sum int
    for x != 0 {
        sum += x % 10
        x /= 10
    }
    for y != 0 {
        sum += y % 10
        y /= 10
    }
    return sum
}
```