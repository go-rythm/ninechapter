package bfs4

import (
	"strconv"
	"strings"
)

// https://www.lintcode.com/problem/7/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := []*TreeNode{root}
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if node == nil {
			continue
		}
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	for queue[len(queue)-1] == nil {
		queue = queue[:len(queue)-1]
	}

	sb := &strings.Builder{}
	sb.WriteByte('{')
	sb.WriteString(strconv.Itoa(queue[0].Val))
	for i := 1; i < len(queue); i++ {
		if queue[i] == nil {
			sb.WriteString(",#")
		} else {
			sb.WriteByte(',')
			sb.WriteString(strconv.Itoa(queue[i].Val))
		}
	}
	sb.WriteByte('}')
	return sb.String()
}
func Deserialize(data string) *TreeNode {
	if data == "{}" {
		return nil
	}
	vals := strings.Split(data[1:len(data)-1], ",")
	queue := []*TreeNode{}
	val, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: val}
	queue = append(queue, root)
	index := 0
	isLeftChild := true

	for i := 1; i < len(vals); i++ {
		if vals[i] != "#" {
			val, _ := strconv.Atoi(vals[i])
			node := &TreeNode{Val: val}
			if isLeftChild {
				queue[index].Left = node
			} else {
				queue[index].Right = node
			}
			queue = append(queue, node)
		}
		if !isLeftChild {
			index++
		}
		isLeftChild = !isLeftChild
	}
	return root
}
