package dfs

import "math"

type Result2 struct {
	minCost int
}

func NewResult2() *Result {
	return &Result{minCost: 1000000}
}

func minCost2(n int, roads [][]int) int {
	graph := constructGraph(roads, n)
	visited := make(map[int]bool)
	path := []int{1}
	result := NewResult()
	visited[1] = true

	travel2(1, n, path, visited, 0, graph, result)

	return result.minCost
}

func travel2(city int, n int, path []int, visited map[int]bool, cost int, graph [][]int, result *Result) {
	if len(visited) == n {
		result.minCost = int(math.Min(float64(result.minCost), float64(cost)))
		return
	}

	for i := 1; i < len(graph[city]); i++ {
		if visited[i] {
			continue
		}
		if hasBetterPath(graph, path, i) {
			continue
		}
		visited[i] = true
		path = append(path, i)
		travel2(i, n, path, visited, cost+graph[city][i], graph, result)
		path = path[:len(path)-1]
		visited[i] = false
	}
}

func constructGraph2(roads [][]int, n int) [][]int {
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
		for j := range graph[i] {
			graph[i][j] = 100000
		}
	}

	for _, road := range roads {
		a, b, c := road[0], road[1], road[2]
		if graph[a][b] > c {
			graph[a][b] = c
		}
		if graph[b][a] > c {
			graph[b][a] = c
		}
	}

	return graph
}

func hasBetterPath(graph [][]int, path []int, city int) bool {
	pathLength := len(path)
	for i := 1; i < pathLength; i++ {
		path_i_1 := path[i-1]
		path_i := path[i]
		path_last := path[pathLength-1]

		if graph[path_i_1][path_i]+graph[path_last][city] >
			graph[path_i_1][path_last]+graph[path_i][city] {
			return true
		}
	}
	return false
}
