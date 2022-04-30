package 树

/**
题目：
给定一个 N 叉树，返回其节点值的层序遍历，每一层单独放在一个切片中，返回一个二维切片。（即从左到右，逐层遍历）。

思路：
BFS T = S = O(n)
*/
//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (ans [][]int) {
	//特殊情况：树为空，返回空
	if root == nil {
		return
	}
	q := []*Node{root}
	for q != nil {
		level := []int{}
		tmp := q //tmp保存的是当前待处理的这一层的所有节点
		q = nil
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...) //将当前层的所有孩子节点，都重新存入队列q中
		}
		ans = append(ans, level)
	}
	return
}
