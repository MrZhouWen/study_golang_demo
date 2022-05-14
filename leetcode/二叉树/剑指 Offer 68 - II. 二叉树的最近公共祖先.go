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
	//root也可以作为自己的祖先
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q) //找寻，p或者q是否有一个、或者都在left子树中
	right := lowestCommonAncestor(root.Right, p, q)

	//说明p, q 分别在 root 的左右子树中，所以root为最近公共祖先
	if left != nil && right != nil {
		return root
	}

	//说明right == nil ,都在左子树上，最后返回的left就是最近祖先
	if left != nil {
		return left
	}

	return right
}
