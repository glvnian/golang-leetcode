package day_08

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serializeV1(root *TreeNode) string {
	var res []string
	var stack []*TreeNode
	if root == nil {
		return ""
	}
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, fmt.Sprintf("%d", cur.Val))
			stack = append(stack, cur)
			cur = cur.Left
		}

		if cur == nil {
			res = append(res, "#")
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Right
	}

	return strings.Join(res, ",")
}

//
//// v1
//// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	node := root
	if node == nil {
		return "#"
	}
	leftStr := this.serialize(node.Left)
	rightStr := this.serialize(node.Right)
	return fmt.Sprintf("%d,%s,%s", node.Val, leftStr, rightStr)
}

//
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var dataList = strings.Split(data, ",")
	var r = []int{0}
	return dfs(dataList, r)
}

func dfs(dataList []string, i []int) *TreeNode {
	if i[0] == len(dataList) {
		return nil
	}
	str := dataList[i[0]]
	i[0]++
	if str == "#" {
		return nil
	}

	num, _ := strconv.Atoi(str)
	node := &TreeNode{Val: num}
	node.Left = dfs(dataList, i)
	node.Right = dfs(dataList, i)
	return node
}
