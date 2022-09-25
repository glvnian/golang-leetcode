package day_10

func minimumLengthEncoding(words []string) int {
	var root = buildTrie(words)
	var size = []int{0}
	trieNodeSize(root, size, 1)
	return size[0]
}

func buildTrie(words []string) *Trie {
	var root, node *Trie
	root = &Trie{}

	for _, word := range words {
		node = root
		for i := len(word) - 1; i >= 0; i-- {
			if node.children[word[i]-'a'] == nil {
				node.children[word[i]-'a'] = &Trie{}
			}
			node = node.children[word[i]-'a']
		}
		node.isWord = true
	}

	return root

}

func trieNodeSize(root *Trie, size []int, length int) {
	var ifLeaf = true
	for _, node := range root.children {
		if node != nil {
			ifLeaf = false
			trieNodeSize(node, size, length+1)
		}
	}
	if ifLeaf {
		size[0] = size[0] + length
	}
}
