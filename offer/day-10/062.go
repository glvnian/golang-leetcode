package day_10

import "fmt"

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, ch := range word {
		if this.children[ch-'a'] == nil {
			this.children[ch-'a'] = &Trie{}
			fmt.Println("build tree", ch-'a')
		}
		this = this.children[ch-'a']
	}
	this.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, ch := range word {
		if this.children[ch-'a'] == nil {
			return false
		}
		this = this.children[ch-'a']
	}

	return this.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, ch := range prefix {
		if this.children[ch-'a'] == nil {
			return false
		}
		this = this.children[ch-'a']

	}

	return true
}
