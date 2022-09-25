package day_10

type MagicDictionary struct {
	root *Trie
}

/** Initialize your data structure here. */
func MConstructor() MagicDictionary {
	return MagicDictionary{
		root: &Trie{},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		var node = this.root
		for _, ch := range word {
			if node.children[ch-'a'] == nil {
				node.children[ch-'a'] = &Trie{}
			}
			node = node.children[ch-'a']
		}
		node.isWord = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	var edit, start int
	var node = this.root
	return this.search(node, searchWord, start, edit)
}

func (this *MagicDictionary) search(node *Trie, searchWord string, start, edit int) bool {
	if node == nil {
		return false
	}

	// 如果达到对应字符串最后一个字符对应的节点时该节点的isWord字符串为true，
	// 并且此时正好修改了一个字符串。那么就是需要找的字符
	if node.isWord && start == len(searchWord) && edit == 1 {
		return true
	}

	if start < len(searchWord) && edit <= 1 {
		var found bool
		var i uint8
		for i = 0; i < 26 && !found; i++ {
			if node.children[i] == nil {
				continue
			}
			// 如果到达的节点和字符串中的字符不匹配，则表示修改字符串中的一个字符以匹配前缀树中的路径
			if i != ([]byte(searchWord)[start] - 'a') {
				edit = edit + 1
			}

			found = this.search(node.children[i], searchWord, start+1, edit)
		}
		return found
	}
	return false
}
