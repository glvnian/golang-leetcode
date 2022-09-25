package day_05

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var s1, t1 string
	var b bool
	s1 = "anagram"
	t1 = "nagarama"
	b = isAnagram(s1, t1)
	fmt.Println(b)
}
