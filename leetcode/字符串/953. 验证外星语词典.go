package 字符串

/**
某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。

思路：直接遍历 T= O(字符串数组长度 * 每个字符串的平均长度) S= O(26)
*/
func isAlienSorted(words []string, order string) bool {
	index := [26]int{}
	for i, c := range order {
		index[c-'a'] = i
	}

next: //这里next很精髓
	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			pre, cur := index[words[i-1][j]-'a'], index[words[i][j]-'a']
			if pre > cur {
				return false
			}
			//因为按照字典序排序，前面的高位，只要小于 后面的高位，那么就可以判定，前者小于后者，后面的部分不需要再比对了
			if pre < cur {
				continue next
			}
		}
		if len(words[i-1]) > len(words[i]) {
			return false
		}
	}
	return true
}
