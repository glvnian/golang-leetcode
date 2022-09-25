package day_05

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var strs []string
	var r [][]string
	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	r = groupAnagrams(strs)
	fmt.Println(r)
}
