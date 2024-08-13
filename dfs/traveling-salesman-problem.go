package dfs

import (
	"math"
)

// Result 用于存储最小费用
type Result struct {
	minCost int
}

// NewResult 创建一个新的 Result 实例，并初始化 minCost 为一个较大值
func NewResult() *Result {
	return &Result{
		minCost: math.MaxInt32,
	}
}

// minCost 计算访问所有城市的最小费用
func minCost(n int, roads [][]int) int {
	graph := constructGraph(roads, n)
	visited := make(map[int]bool)
	result := NewResult()
	visited[1] = true

	travel(1, n, visited, 0, graph, result)

	return result.minCost
}

// dfs 深度优先搜索遍历城市
func travel(
	city int,
	n int,
	visited map[int]bool,
	cost int,
	graph [][]int,
	result *Result) {

	// 如果所有城市都已访问，更新最小费用
	if len(visited) == n {
		if cost < result.minCost {
			result.minCost = cost
		}
		return
	}

	for i := 1; i < len(graph[city]); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		travel(i, n, visited, cost+graph[city][i], graph, result)
		delete(visited, i)
	}
}

// constructGraph 构造图的邻接矩阵
func constructGraph(roads [][]int, n int) [][]int {
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
		for j := range graph[i] {
			graph[i][j] = math.MaxInt32
		}
	}

	for _, road := range roads {
		a, b, c := road[0], road[1], road[2]
		graph[a][b] = min(graph[a][b], c)
		graph[b][a] = min(graph[b][a], c)
	}

	return graph
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
