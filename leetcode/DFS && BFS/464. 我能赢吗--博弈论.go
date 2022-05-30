package DFS____BFS

/**
在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和 达到或超过 100 的玩家，即为胜者。
如果我们将游戏规则改为 “玩家 不能 重复使用整数” 呢？
例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
给定两个整数maxChoosableInteger（整数池中可选择的最大数）和desiredTotal（累计和），
若先出手的玩家是否能稳赢则返回 true，否则返回 false 。假设两位玩家游戏时都表现 最佳 。

思路：记忆化搜索 + dfs
T =O(2^n×n) S = O(2^n)
考虑边界情况，当所有数字选完仍无法到达 desiredTotal 时，两人都无法获胜，返回 false。
当所有数字的和大于等于 desiredTotal 时，其中一方能获得胜利，需要通过搜索来判断获胜方。
在游戏中途，假设已经被使用的数字的集合为 usedNumbers，这些数字的和为 currentTotal。
	当某方行动时，如果他能在未选择的数字中选出一个 i，使得 i+currentTotal ≥ desiredTotal，则他能获胜。否则，需要继续通过搜索来判断获胜方。
	在剩下的数字中，如果他能选择一个 i，使得对方在接下来的局面中无法获胜，则他会获胜。否则，他会失败。

根据这个思想设计搜索函数 dfs，其中 usedNumbers 可以用一个整数来表示，从低位到高位，
	第 i 位为 1 则表示数字 i 已经被使用，为 0 则表示数字 i 未被使用。
	如果当前玩家获胜，则返回 true，否则返回 false。
为了避免重复计算，需要使用记忆化的操作来降低时间复杂度。
*/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (maxChoosableInteger+1)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	dp := make([]int8, 1<<maxChoosableInteger) //这里题目说明，maxChoosableInteger <= 20所以，int8即可
	for i := range dp {
		dp[i] = -1 //初始化，表示i的二进制表示的数，比如 1010表示，目前选择了数字 2 和 4，是否能是的先手的玩家获胜
	}

	var dfs func(int, int) int8 //返回0 或者 1
	dfs = func(usedNum, curTotal int) (res int8) {
		dv := &dp[usedNum]
		if *dv != -1 { //说明usedNum的二进制表示的，为1的位置上得数已被选择，且遍历过了,结果是否满足，直接返回即可
			return *dv
		}
		defer func() {
			*dv = res
		}()
		//这里i是下标，对应 i = 0, 对应数字 1, 因为二进制的最低位，是 第一位2^0, 第二位2^1, 第三位2^2....
		//所以数字(i+1)是否被选用，只要检测usedNum的二进制表示的i位是否为1即可
		for i := 0; i < maxChoosableInteger; i++ {
			//表示usedNum 不包含 (i+1), 并且【本次选择(i+1)后，就可胜利 || 或本次选择(i+1)后，对手在此基础上继续游戏会失败，那么也是我赢】
			if usedNum>>i&1 == 0 && (curTotal+i+1 >= desiredTotal || dfs(usedNum|1<<i, curTotal+i+1) == 0) {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0) == 1
}

/**
292. Nim 游戏
你和你的朋友，两个人一起玩Nim 游戏：
	桌子上有一堆石头。
	你们轮流进行自己的回合，你作为先手。
	每一回合，轮到的人拿掉1 - 3 块石头。
	拿掉最后一块石头的人就是获胜者。
假设你们每一步都是最优解。请编写一个函数，来判断你是否可以在给定石头数量为 n 的情况下赢得游戏。如果可以赢，返回 true；否则，返回 false 。

数学：
如果总的石头数目为 4 的倍数时，因为无论你取多少石头，对方总有对应的取法，让剩余的石头的数目继续为 4 的倍数。
对于你或者你的对手取石头时，显然最优的选择是当前己方取完石头后，让剩余的石头的数目为 4 的倍数。
假设当前的石头数目为 x，如果 x 为 4 的倍数时，则此时你必然会输掉游戏；
如果 x 不为 4 的倍数时，则此时你只需要取走 xmod4 个石头时，则剩余的石头数目必然为 4 的倍数，从而对手会输掉游戏。
*/
func canWinNim(n int) bool {
	return n%4 != 0
}
