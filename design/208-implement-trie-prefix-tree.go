/*
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/
 * 2020.11.25
 * Golang 72ms(56.53%), 17.9MB(28.78%)
 */
package design

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	return Trie{isEnd: false, next: [26]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, n := range word {
		if node.next[n-'a'] == nil {
			node.next[n-'a'] = new(Trie)
		}
		node = node.next[n-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, n := range word {
		node = node.next[n-'a']
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, n := range prefix {
		node = node.next[n-'a']
		if node == nil {
			return false
		}
	}
	return true
}
