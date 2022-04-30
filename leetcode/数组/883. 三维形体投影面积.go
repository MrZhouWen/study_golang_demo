package 数组

/**
在n x n的网格grid中，我们放置了一些与 x，y，z 三轴对齐的1 x 1 x 1立方体。
每个值v = grid[i][j]表示 v个正方体叠放在单元格(i, j)上。
现在，我们查看这些立方体在 xy、yz和 zx平面上的投影。
投影就像影子，将 三维 形体映射到一个 二维 平面上。从顶部、前面和侧面看立方体时，我们会看到“影子”。
返回 所有三个投影的总面积 。

思路：一趟遍历 ==> 得出二维数组的每一行 和 每一列 的最大值
T= O(n^2) S= O(1)
*/
func projectionArea(grid [][]int) int {
	var xy, yz, zx int
	for i, row := range grid {
		yzMax, zxMax := 0, 0 //每一列最高，每一行最高
		for j, v := range row {
			if v > 0 {
				xy++
			}
			zxMax = max(zxMax, v)
			yzMax = max(yzMax, grid[j][i]) //找到第i列的最高
		}
		//第i行遍历结束，第i列也遍历结束了
		yz += yzMax
		zx += zxMax
	}
	return xy + yz + zx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
