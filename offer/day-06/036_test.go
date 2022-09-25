package day_06

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	var tokens []string
	var r int
	tokens = []string{"2", "1", "+", "3", "*"}
	tokens = []string{"4", "13", "5", "/", "+"}
	r = evalRPN(tokens)
	fmt.Println(r)
}
