## 剑指 Offer II 115. 重建序列

### 题目传送门

[点击这里](https://leetcode.cn/problems/ur2n8P/)

### 解题思路

这道题不好想到用拓扑排序，所以比较有难度，题意的思路如下，由于`sequences`中的任意一个序列都是`nums`的子序列，所以可以根据任意一个序列中的元素顺序来构建`nums`中元素的有向图，但是如何保证题目所说的`nums`是唯一且超短呢？只需要遍历所有的序列元素，将其都添加则可以保证超短。至于唯一性，在遍历`sequences`中的任意一个序列，构建有向图的过程中，建立入度数组，当入度为0则入队列，说明从该元素位置继续向后，所以此刻如果队列中的元素大于1个，则不满足唯一性，在遍历完后将该元素指向的其他元素的入度-1即可。

### 代码

```go
func sequenceReconstruction(nums []int, sequences [][]int) bool {
    // 拓扑排序
    // 建立有向图，根据题意，sequences中的所有元素都是nums的子序列，可以通过子序列中元素的顺序来表示有向
    n := len(nums)
    g := make([][]int, n+1)
    indegree := make([]int, n+1)
    for _, v := range sequences {
        for i := 1; i < len(v);i ++ {
            x, y := v[i-1], v[i]
            g[x] = append(g[x], y)
            // 入度加1
            indegree[y] ++ 
        }
    }

    q := []int{}
    for i := 1;i <= n;i ++ {
        if indegree[i] == 0 {
            q = append(q, i)
        }
    }
    for len(q) > 0 {
        // 如果同一位置可以插入的元素为多个，则直接返回false
        if len(q) > 1 {
            return false
        }
        cur := q[0]
        q = q[1:]
        // 将cur指向的元素，入度-1
        for _, v := range g[cur] {
            indegree[v] -- 
            if indegree[v] == 0 {
                q = append(q, v)
            }
        } 
    }
    return true
}

```