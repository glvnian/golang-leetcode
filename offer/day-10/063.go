package day_10

import (
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	var sentenceWords = []string{}
	var root *Trie

	root = buildDict(dictionary)

	for _, word := range strings.Split(sentence, " ") {
		var node = root
		sentenceWords = append(sentenceWords, findPrefix(node, word))

	}
	return strings.Join(sentenceWords, " ")
}

func buildDict(dictionary []string) *Trie {
	var root = &Trie{}
	for _, word := range dictionary {
		node := root
		for _, ch := range word {
			if node.children[ch-'a'] == nil {
				node.children[ch-'a'] = &Trie{}
			}
			node = node.children[ch-'a']
		}
		node.isWord = true

	}

	return root
}

func findPrefix(root *Trie, word string) string {
	node := root
	var tWord string
	for i, ch := range word {
		if node.isWord || node.children[ch-'a'] == nil {
			break
		}
		tWord = word[:i+1]
		node = node.children[ch-'a']
	}
	if node.isWord {
		return tWord
	}
	return word
}
