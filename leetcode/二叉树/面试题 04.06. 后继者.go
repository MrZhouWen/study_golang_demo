package 二叉树

/**
请找出二叉搜索树中的指定节点的 直接后继节点（就是中序遍历的后继节点）

思路 ： 中序遍历 T= O(n) S = O(n)
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	st := []*TreeNode{}
	var pre, cur *TreeNode = nil, root
	for len(st) > 0 || cur != nil {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}
	return nil
}
