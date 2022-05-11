package 字符串

import (
	"sort"
	"strings"
	"unicode"
)

/**
给你一个日志数组 logs。每条日志都是以空格分隔的字串，其第一个 字为字母与数字混合的 标识符 。
有两种不同类型的日志：
	字母日志：除标识符之外，所有字均由小写字母组成
	数字日志：除标识符之外，所有字均由数字组成
请按下述规则将日志重新排序：
	所有 字母日志 都排在 数字日志 之前。
	字母日志 在内容不同时，忽略标识符后，按内容字母顺序排序；在内容相同时，按标识符排序。
	数字日志 应该保留原来的相对顺序。
返回日志的最终顺序。

思路：自定义排序 T = O(nlogn)，其中 n 是 logs 的字符数，是平均情况下内置排序的时间复杂度。 S = O(1)

比较时，先将数组日志按照第一个空格分成两部分字符串，其中第一部分为标识符。
第二部分的首字符可以用来判断该日志的类型。两条日志进行比较时，需要先确定待比较的日志的类型，然后按照以下规则进行比较：
	字母日志始终小于数字日志。
	数字日志保留原来的相对顺序。当使用稳定的排序算法时，可以认为所有数字日志大小一样。
	字母日志进行相互比较时，先比较第二部分的大小；如果相等，则比较标识符大小。比较时都使用字符串的比较方式进行比较。
*/
func reorderLogFiles(logs []string) []string {
	//sort.SliceStable是 less func(i, j int) bool{} 返回logs[i] <= logs[j]
	sort.SliceStable(logs, func(i, j int) bool {
		s, t := logs[i], logs[j]
		s1 := strings.SplitN(s, " ", 2)[1] //除去第一个标识符，后面的部分
		s2 := strings.SplitN(t, " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		if isDigit2 && isDigit1 {
			return false //都是数字日志，顺序不变，不存在 logs[i] <= logs[j]的说法
		}
		if !isDigit1 && !isDigit2 {
			return s1 < s2 || s1 == s2 && s < t
		}
		return !isDigit1 //logs[i]是字符串 <= logs[j]是数字日志
	})
	return logs
}
