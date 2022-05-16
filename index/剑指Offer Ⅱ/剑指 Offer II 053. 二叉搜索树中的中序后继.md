## 剑指 Offer II 053. 二叉搜索树中的中序后继

### 题目传送门

[点击这里](https://leetcode.cn/problems/P5rCT8/)

### 解题思路

这道题可以是用最基本的中序遍历的思路，这个应用于所有二叉树，中序遍历可以看一下[94. 二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/)这道题的两种方法，dfs和bfs均可。但这道题说明了二叉树为二叉搜索树，二叉搜索树的性质是中序遍历是递增的数列，所以我们可以利用这一个方法，如果想找到某一个节点的后继值，也就是找其下一个节点，即比其他大的最小节点，根据二叉搜索树的性质，如果该节点的右子树不为空，那么后继节点就在右子树上的最左边节点，如果右子树为空则遍历，遍历的规则如下，如果当前节点值比要找后继节点的节点值p大，说明p的后继是node或者是node的左子树，也就是往小了找，如果小于等于的话，则可能出现在右子树，往大了找。

### 代码

```go
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    // 中序遍历 bfs方法
    q := []*TreeNode{}
    pre := new(TreeNode)
    cur := root
    for len(q) > 0 || cur != nil {
        for cur != nil {
            q = append(q, cur)
            cur = cur.Left
        }
        cur = q[len(q)-1]
        q = q[:len(q)-1]
        if pre == p {
            return cur
        }
        pre = cur
        cur = cur.Right
    }
    return nil
}
```

```go
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    // 这里的声明不能用new，new会声明一个val为0的不为空的指针
    var res *TreeNode
    if p.Right != nil {
        res = p.Right
        for res.Left != nil {
            res = res.Left
        }
        return res
    }
    // 如果右子树为空，则从root开始遍历
    node := root
    for node != nil {
        if node.Val > p.Val {
            res = node
            node = node.Left
        } else {
            node = node.Right
        }
    }
    return res
}
```