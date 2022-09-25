package day_10

import (
	"fmt"
	"testing"
)

func TestMagicDictionary(t *testing.T) {
	var dictionary []string
	var searchWord string
	var r bool
	obj := MConstructor()

	dictionary = []string{"hello", "leetcode"}
	//searchWord = "hello"
	searchWord = "hhllo"
	obj.BuildDict(dictionary)
	r = obj.Search(searchWord)
	fmt.Println(r)
}
