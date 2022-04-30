package DFS____BFS

/**
有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻，类似 「这样」。
“太平洋”处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵heights，heights[r][c]表示坐标 (r, c) 上单元格 高于海平面的高度 。
岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
返回 网格坐标 result的 2D列表 ，其中result[i] = [ri, ci]表示雨水可以从单元格 (ri, ci) 流向 太平洋和大西洋 。

思路：T = O(mn) S = (mn)
方法一：深度优先搜索【反向搜索】
雨水的流动方向是从高到低，每个单元格上的雨水只能流到高度小于等于当前单元格的相邻单元格。
从一个单元格开始，通过搜索的方法模拟雨水的流动，则可以判断雨水是否可以从该单元格流向海洋。
如果直接以每个单元格作为起点模拟雨水的流动，则会重复遍历每个单元格，导致时间复杂度过高。
为了降低时间复杂度，可以从矩阵的边界开始反向搜索寻找雨水流向边界的单元格，反向搜索时，每次只能移动到高度相同或更大的单元格。
由于矩阵的左边界和上边界是太平洋，矩阵的右边界和下边界是大西洋，因此从矩阵的左边界和上边界开始反向搜索即可找到雨水流向太平洋的单元格，
从矩阵的右边界和下边界开始反向搜索即可找到雨水流向大西洋的单元格。
可以使用深度优先搜索实现反向搜索，搜索过程中需要记录每个单元格是否可以从太平洋反向到达以及是否可以从大西洋反向到达。
反向搜索结束之后，遍历每个网格，如果一个网格既可以从太平洋反向到达也可以从大西洋反向到达，则该网格满足太平洋和大西洋都可以到达，将该网格添加到答案中。
*/
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		for _, d := range dirs {
			if nx, ny := x+d.x, y+d.y; 0 <= nx && nx < m && 0 <= ny && ny < n && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for j := 1; j < n; j++ {
		dfs(0, j, pacific)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	for j := 0; j < n-1; j++ {
		dfs(m-1, j, atlantic)
	}

	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}

/** BFS
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {
    m, n := len(heights), len(heights[0])
    pacific := make([][]bool, m)
    atlantic := make([][]bool, m)
    for i := range pacific {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }

    bfs := func(x, y int, ocean [][]bool) {
        if ocean[x][y] {
            return
        }
        ocean[x][y] = true
        q := []pair{{x, y}}
        for len(q) > 0 {
            p := q[0]
            q = q[1:]
            for _, d := range dirs {
                if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !ocean[x][y] && heights[x][y] >= heights[p.x][p.y] {
                    ocean[x][y] = true
                    q = append(q, pair{x, y})
                }
            }
        }
    }
    for i := 0; i < m; i++ {
        bfs(i, 0, pacific)
    }
    for j := 1; j < n; j++ {
        bfs(0, j, pacific)
    }
    for i := 0; i < m; i++ {
        bfs(i, n-1, atlantic)
    }
    for j := 0; j < n-1; j++ {
        bfs(m-1, j, atlantic)
    }

    for i, row := range pacific {
        for j, ok := range row {
            if ok && atlantic[i][j] {
                ans = append(ans, []int{i, j})
            }
        }
    }
    return
}
*/
