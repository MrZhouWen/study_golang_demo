package 哈希表

import "unicode"

/**
题目：
给定一个段落 (paragraph) 和一个禁用单词列表 (banned)。返回出现次数最多，同时不在禁用列表中的单词。
题目保证至少有一个词不在禁用列表中，而且答案唯一。
禁用列表中的单词用小写字母表示，不含标点符号。段落中的单词不区分大小写。答案都是小写字母。

思路：
为了判断给定段落中的每个单词是否在禁用单词列表中，需要使用哈希集合存储禁用单词列表中的单词。以下将禁用单词列表中的单词称为禁用单词。
遍历段落 paragraph，得到段落中的所有单词，并对每个单词计数，使用哈希表记录每个单词的计数。由于每个单词由连续的字母组成，因此当遇到一个非字母的字符且该字符的前一个字符是字母时，即为一个单词的结束，如果该单词不是禁用单词，则将该单词的计数加 11。如果段落的最后一个字符是字母，则当遍历结束时需要对段落中的最后一个单词判断是否为禁用单词，如果不是禁用单词则将次数加 11。
在遍历段落的过程中，对于每个单词都会更新计数，因此遍历结束之后即可得到最大计数，即出现次数最多的单词的出现次数。
遍历段落之后，遍历哈希表，寻找出现次数等于最大计数的单词，该单词即为最常见的单词。
*/
func mostCommonWord(paragraph string, banned []string) string {
	ban := map[string]bool{} //保存禁止的词
	for _, s := range banned {
		ban[s] = true
	}
	freq := map[string]int{} //记录每一个单词出现的次数
	word := []byte{}
	maxFreq := 0
	for i, n := 0, len(paragraph); i <= n; i++ { //这里 i=n 很精髓啊，因为最后一个单词遍历完要处理一下
		if i < n && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if word != nil { //遍历完一个整的单词
			s := string(word)
			if !ban[s] {
				freq[s]++
				maxFreq = max(maxFreq, freq[s])
			}
			word = nil
		}
	}

	//遍历freq,寻找出现次数 = maxFreq 的单词
	for str, num := range freq {
		if num == maxFreq {
			return str
		}
	}
	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
