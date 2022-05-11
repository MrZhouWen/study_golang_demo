package 二叉树

/**
题目：
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

思路：
递归
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//说明root的左右孩子分别是p q
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
