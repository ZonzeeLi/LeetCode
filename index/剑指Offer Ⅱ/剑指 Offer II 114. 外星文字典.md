## 剑指 Offer II 114. 外星文字典

### 题目传送门

[点击这里](https://leetcode.cn/problems/Jf1JuT/)

### 解题思路

这是一道我认为很有难度值得思考的题，首先题意比较好分析，就是根据给定的按照某一种字典序排列好的内容words，来返回满足条件的一种字典序。题意我们大致了解了，这道题使用的方法是拓扑排序+dfs/bfs，整体思路是使用有向图，我们可以比较相邻的两个字符串，然后再遍历字符直到第一次遇到不相同的字符，然后就可以判断出两种字符的字典序，做一个有向连接，这是dfs和bfs的初始化，bfs的初始化还多一个入度的map记录。接下来分成dfs和bfs两种方法来解释，首先对于dfs，dfs代码的思路是对每一个未添加到结果的字符进行深度遍历，如果遍历某个字符的后面连接其他字符，这些其他字符也未遍历过，就继续向下，如果是遍历过的，那么说明其已经计算过，添加到结果了，但如果也是在遍历中的状态，说明从某一字符开始遍历后续字符，注意这里这个字符并没有遍历完成，记录的状态还是遍历中，然后后续的某一字符开始继续遍历，当它遍历回到前面的这个字符也是遍历中的状态，说明有向图中形成了环，那么就没有合法的字典序。bfs的思路基本和dfs类似，只不过bfs的思路更好理解，直接用有向入度来做，这个需要有拓扑排序和图论的基本知识，能好理解一些，这里不赘述图论的基本知识了，当我们从一个入度为0的字符开始向后遍历，然后将该字符和它指向的字符之间的有向连接断开，如果指向的字符在这一次变化中入度变为0，则入队列，可以之后从其开始遍历，bfs的整体思路是这样，不过这里的go语言代码利用了slice的底层数组，写的非常优雅，这里也要学会和理解go语言slice的底层原理才能理解好这段代码，为什么感觉没有对order操作就改变了order的数组值。

### 代码

```go
func alienOrder(words []string) string {
	g := map[byte][]byte{}
	// 初始化
	for _, c := range words[0] {
		g[byte(c)] = nil
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			// 初始化
			g[byte(c)] = g[byte(c)]
		}
		// fmt.Println(g)
		for j := 0; j < len(s) && j < len(t); j++ {
			// 遍历前后字符串，当出现不同字符则，制作一个有向连接
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				continue next
			}
		}
		// 如果类似"ab"，"a" 这种情况，不合法。
		if len(s) > len(t) {
			return ""
		}
	}
	// 未访问，访问中，访问结束三种种状态
	const visiting, visited = 1, 2
	order := make([]byte, len(g))
	i := len(g) - 1
	state := map[byte]int{}
	// dfs深度遍历
	var dfs func(u byte) bool
	dfs = func(u byte) bool {
		// 记录当前字符为遍历状态
		state[u] = visiting
		// 遍历当前字符后的字符情况
		for _, v := range g[u] {
			// 如果后面的字符处于未访问，要继续dfs，否则如果后面的字符处于访问中，在遍历的过程中又遍历回去，那么说明有向图中存在环，直接返回false
			if state[v] == 0 {
				if !dfs(v) {
					return false
				}
			} else if state[v] == visiting {
				return false
			}
		}
		// 如果在map中存放的情况为其后面没有字符，或者后面的字符已经访问过，说明已经添加过了，那么就可以存放在order中，从后开始存放
		order[i] = u
		i--
		// 记录为访问过的状态
		state[u] = visited
		return true
	}
	for u := range g {
		// 如果u没有访问过，开始dfs
		if state[u] == 0 && !dfs(u) {
			return ""
		}
	}
	return string(order)
}
```

```go
func alienOrder(words []string) string {
	g := map[byte][]byte{}
	// 记录入度
	inDeg := map[byte]int{}
	for _, c := range words[0] {
		inDeg[byte(c)] = 0
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		// 初始化
		for _, c := range t {
			inDeg[byte(c)] += 0
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			// 同上面的思路一样，只不过入度map也同时要更新
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				inDeg[t[j]]++
				continue next
			}
		}
		if len(s) > len(t) {
			return ""
		}
	}

	order := make([]byte, len(inDeg))
	// 这个q的底层数组为order
	q := order[:0]
	for u, d := range inDeg {
		// 如果入度为0，则入栈，说明可以作为第一个字典序的字符
		if d == 0 {
			q = append(q, u)
		}
	}
	// bfs
	for len(q) > 0 {
		u := q[0]
		// 这里q开始只想order[1:]，底层数组是go语言的slice基本知识，建议了解一下
		q = q[1:]
		// 这里可以打印一下q和order
		for _, v := range g[u] {
			// 遍历其后面的字符，入度-1相当于在有向图的连接上让u和v的连接断开。如果v的入度变为0，则入队列
			if inDeg[v]--; inDeg[v] == 0 {
				// 会改变order
				q = append(q, v)
			}
			// 这里可以打印一下q和order
		}
	}
	if cap(q) == 0 {
		return string(order)
	}
	return ""
}
```
