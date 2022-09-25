package day_10

type Trie struct {
	isWord   bool
	children [26]*Trie
}
