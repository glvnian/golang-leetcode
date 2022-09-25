package day_03

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	var s, p, r string
	s = "ADOBECODEBANC"
	p = "ABC"

	//s = "ab"
	//p = "b"

	//s = ""
	//p = "b"
	//r = minWindow(s, p)
	r = minWindowV1(s, p)
	fmt.Println(r)

}
