package 字符串

/**
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。
给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

思路：分类讨论 || 双指针
假设 m = len(first), n = len(second)
	往 first 中插入一个字符得到 second，此时 n−m=1，second 比 first 多一个字符，其余字符都相同；
	从 first 中删除一个字符得到 second，此时 m−n=1，first 比 second 多一个字符，其余字符都相同；
	将 first 中的一个字符替换成不同的字符得到 second，此时 m=n，first 和 second 恰好有一个字符不同。

T = O(m) S = O(1)
*/
func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n { //前两种情况统一处理，将较长的字符串放到first
		return oneEditAway(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i, ch := range second {
		if first[i] != byte(ch) {
			if m > n {
				return first[i+1:] == second[i:]
			} else {
				return first[i+1:] == second[i+1:]
			}
		}
	}
	return true
}
