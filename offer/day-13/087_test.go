package day_13

import (
	"fmt"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	s := "25525511135"
	s = "010010"
	r := restoreIpAddresses(s)
	fmt.Println(r)
}

func TestNum(t *testing.T) {
	a := "1110"
	fmt.Println(isIpNode(a))
}
