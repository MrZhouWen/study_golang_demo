package DFS____BFS

/*
题目：
树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
给你一棵包含n个节点的树，标记为0到n - 1 。给定数字n和一个有 n - 1 条无向边的 edges列表（每一个边都是一对标签），
其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。

请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。

树的 高度 是指  根节点  和  叶子节点  之间最长向下路径上  边  的数量。

思路：
广度优先搜索 T(n) S(n)
只需要找到树中距离最远的2各节点x, y即可，那么所求节点，在x, y之间的路径上
*/
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	parents := make([]int, n)
	bfs := func(start int) (x int) {
		vis := make([]bool, n)
		vis[start] = true
		q := []int{start}
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, y := range g[x] {
				if !vis[y] {
					vis[y] = true
					parents[y] = x
					q = append(q, y)
				}
			}
		}
		return
	}

	x := bfs(0) //首先任意的从0出发(因为有可能只有一个0节点，所以不失一般性的选0)，找到距离最远的节点x
	y := bfs(x) //找到距离x最远的节点y

	path := []int{} //这表示距离最远的2个节点的路径
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}
