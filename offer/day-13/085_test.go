package day_13

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 2
	result := generateParenthesis(n)
	for _, item := range result {
		fmt.Println(item)
	}
}

func setArray(a []string) {
	a = append(a, "1")
	a = append(a, "2")
}

func TestSetArray(t *testing.T) {
	r := []string{"0"}
	setArray(r)
	fmt.Println(r)
}
