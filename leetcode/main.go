package main

/*
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

	path := make([]int, n) //这表示距离最远的2个节点的路径
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
