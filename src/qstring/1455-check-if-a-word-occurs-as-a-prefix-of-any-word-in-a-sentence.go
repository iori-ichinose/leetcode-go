/*
 * https://leetcode-cn.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
 * 2020.11.23
 * Golang 0ms(100.00%), 1.9MB(77.97%)
 */
package qstring

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, n := range words {
		if strings.HasPrefix(n, searchWord) {
			return i + 1
		}
	}
	return -1
}

func IsPrefixOfWordTb() {
	print("Testbench isPrefixOfWord: ")
	println(isPrefixOfWord("i love eating burger", "burg"))
}
