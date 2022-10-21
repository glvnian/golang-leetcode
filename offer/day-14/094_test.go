package day_14

import (
	"fmt"
	"testing"
)

func TestMinCut(t *testing.T) {
	s := "acb"
	r := minCut(s)
	fmt.Println(r)
}
