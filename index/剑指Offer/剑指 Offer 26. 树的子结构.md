## 剑指 Offer 26. 树的子结构

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)

### 解题思路

这道题可以用递归的思想来做，如果B是A的子结构，那么B和A的某一个子结构完全一致或者和A相同，所以这里要判断的是`isSame(A, B)`和`isSubStructure(A.Left, B)`和`isSubStructure(A.Right, B)`的任意一个为true即可，这样就递归了下去，而在判断两个子结构是否相同时，要一直判断`A`和`B`的值还有`A.Left`和`B.Left`，`A.Right`和`B.Right`，另外要注意的是，在`isSame()`的判断中，如果B判断到空了，说明B的此分支已经完全匹配，可以直接返回true，但如果A不是空，那么B还存在节点，就说明B不是A的子结构。

### 代码

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return isSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil {
		return false
	}

	return A.Val == B.Val && isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}

```