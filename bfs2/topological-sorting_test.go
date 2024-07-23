package bfs2

import (
	"encoding/json"
	"testing"
)

func TestTopSort(t *testing.T) {
	type args struct {
		graph []*DirectedGraphNode
	}
	tests := []struct {
		name string
		args args
		want []*DirectedGraphNode
	}{
		{
			name: "test1",
			args: args{
				graph: func() []*DirectedGraphNode {
					node0 := &DirectedGraphNode{Label: 0}
					node1 := &DirectedGraphNode{Label: 1}
					node2 := &DirectedGraphNode{Label: 2}
					node3 := &DirectedGraphNode{Label: 3}
					node4 := &DirectedGraphNode{Label: 4}
					node5 := &DirectedGraphNode{Label: 5}
					node0.Neighbors = []*DirectedGraphNode{node1, node2, node3}
					node1.Neighbors = []*DirectedGraphNode{node4}
					node2.Neighbors = []*DirectedGraphNode{node4, node5}
					node3.Neighbors = []*DirectedGraphNode{node4, node5}
					return []*DirectedGraphNode{node0, node1, node2, node3, node4, node5}
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopSort(tt.args.graph)
			bytes, _ := json.MarshalIndent(got, "", "\t")
			t.Log(string(bytes))
		})
	}
}
