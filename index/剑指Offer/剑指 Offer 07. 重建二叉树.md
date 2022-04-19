## 剑指 Offer 07. 重建二叉树

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)

### 解题思路

这道题首先要知道前序遍历和中序遍历分别是什么，这里简单说一下前序遍历和中序遍历的规则，可以这样子记，前序遍历`[根，左，右]`，中序遍历`[左，根，右]`，后序遍历`[左，右，根]`，这里利用了题干中不含重复数字，我们可以确定好根节点的位置，然后就可以在前序遍历和中序遍历的两个数组中通过比对得到根节点的位置，这样就能把左右子树的前序遍历和中序遍历得到，以此方法递归即可。

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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    // 整体是递归思路
    root := &TreeNode{preorder[0], nil, nil}
    
    // 前序遍历[根，左，右]，中序遍历[左，根，右]
    // 每一个左子树和右子树可能继续向下延申，利用题意不含重复数字，我们确定好根节点的位置，这样子能够得到左右子树的前序和中序遍历了。然后进行递归即可
    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    root.Left = buildTree(preorder[1:i+1], inorder[:i]) // 得到左子树的前序遍历和中序遍历
    root.Right = buildTree(preorder[i+1:], inorder[i+1:]) // 得到右子树的前序遍历和中序遍历
    return root
}
```