package DFS____BFS

import "sort"

/**
题目：https://leetcode.cn/problems/cut-off-trees-for-golf-event/

思路 ： BFS
T = O(m^2 * n^2), 最多m*n个树，计算每2个树之间的最短距离，时间m*n
S = O(m*n)
*/
var dirs = []struct{ x, y int }{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func cutOffTree(forest [][]int) (ans int) {
	type pair struct{ dis, x, y int }
	trees := []pair{}
	for i, t := range forest {
		for j, d := range t {
			if d > 1 { //只保存有树的
				trees = append(trees, pair{d, i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].dis < trees[j].dis
	})
	//bfs 每一轮只计算 2个点之间的最短路径
	bfs := func(sx, sy, tx, ty int) int {
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m) //记录每一个点是否被访问过 或者 已经在队列中了
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []pair{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, dir := range dirs {
				if x, y := p.x+dir.x, p.y+dir.y; x >= 0 && x < m && y >= 0 && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, pair{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}
	//起点
	preX, preY := 0, 0
	for _, t := range trees {
		dis := bfs(preX, preY, t.x, t.y)
		if dis < 0 {
			return -1
		}
		ans += dis
		preX, preY = t.x, t.y
	}
	return
}
