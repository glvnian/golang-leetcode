package day_10

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var r bool
	var word, prefix string

	//["Trie","insert","search","search","startsWith","insert","search"]
	//	[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]

	obj := Constructor()

	word = "apple"
	obj.Insert(word)
	r = obj.Search(word)
	fmt.Println("search: ", word, r)
	prefix = "app"
	r = obj.StartsWith(prefix)
	fmt.Println("StartsWith: ", prefix, r)

}
