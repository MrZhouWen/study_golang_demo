package 字符串

import "strings"

/**
请你将句子转换为 “山羊拉丁文（Goat Latin）”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。山羊拉丁文的规则如下：
	1、如果单词以元音开头（'a', 'e', 'i', 'o', 'u'），在单词后添加"ma"。例如，单词 "apple" 变为 "applema" 。
	2、如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。例如，单词 "goat" 变为 "oatgma" 。
	3、根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从 1 开始。例如，在第一个单词后加 "a" ，在第二个后添加 "aa" ，以此类推。
返回将 sentence 转换为山羊拉丁文后的句子。

输入：sentence = "I speak Goat Latin"
输出："Imaa peaksmaaa oatGmaaaa atinLmaaaaa"

思路：模拟
*/
var vowels = map[byte]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func toGoatLatin(sentence string) string {
	ans := &strings.Builder{}
	for i, cnt, n := 0, 1, len(sentence); i < n; i++ { //cnt表示单词的索引，即当前是第几个单词
		if cnt > 1 { //到一个新单词了，插入一个空格
			ans.WriteByte(' ')
		}
		start := i
		for ; i < n && sentence[i] != ' '; i++ {
		} //遍历到这个单词末尾的空格结束，此时i指向空格' '
		if _, ok := vowels[sentence[start]]; ok {
			ans.WriteString(sentence[start:i])
		} else {
			ans.WriteString(sentence[start+1 : i])
			ans.WriteByte(sentence[start])
		}
		ans.WriteString("ma")
		ans.WriteString(strings.Repeat("a", cnt))
		cnt++ //即将到下一个单词
	}
	return ans.String()
}
