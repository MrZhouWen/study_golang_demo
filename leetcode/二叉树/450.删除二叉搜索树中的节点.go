package 二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
删除二叉搜索树中的指定节点，使之还是一颗二叉搜索树, 并返回root节点

递归 T= O(n) S=O(n)
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	default: //root.Val == key
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left //寻找root的直接后继顶替root的位置
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
