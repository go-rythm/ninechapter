package bfs4

import "testing"

func Test_serialize(t *testing.T) {
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 20}
	node6 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node6
	node3.Right = node7
	got := Serialize(node1)
	got2 := Serialize(Deserialize(got))
	t.Log(got == got2)
}
