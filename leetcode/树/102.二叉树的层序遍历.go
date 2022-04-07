package 树

/**
题目：
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

思路：
 * 广度优先遍历
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			ret[i] = append(ret[i], q[j].Val)
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}

	return ret
}
