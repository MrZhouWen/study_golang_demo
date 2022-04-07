package 字符串

import "strings"

/*
题目：
给定两个字符串, s和goal。如果在若干次旋转操作之后，s能变成goal，那么返回true。
s的 旋转操作 就是将s 最左边的字符移动到最右边。
例如, 若s = 'abcde'，在旋转一次之后结果就是'bcdea'。

思路：
strings.Contains T = S = O(n)
因为s+s包含了所有s旋转后的情况，在s+s中搜索goal，十分巧妙
*/
func rotateString(s, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
