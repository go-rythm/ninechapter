package dfs

func minCost3(n int, roads [][]int) int {
	inf := 1000000000
	graph := constructGraph(roads, n)
	stateSize := 1 << n
	f := make([][]int, stateSize)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[1][0] = 0 // 初始状态，城市1（即城市0）出发

	for state := 0; state < stateSize; state++ {
		for i := 1; i < n; i++ { // 从城市1（即城市0）以外的其他城市出发
			if state&(1<<i) == 0 {
				continue
			}
			prevState := state ^ (1 << i)
			for j := 0; j < n; j++ {
				if prevState&(1<<j) == 0 {
					continue
				}
				f[state][i] = min(f[state][i], f[prevState][j]+graph[j][i])
			}
		}
	}

	minimalCost := inf
	for i := 0; i < n; i++ {
		minimalCost = min(minimalCost, f[stateSize-1][i])
	}

	return minimalCost
}

func constructGraph3(roads [][]int, n int) [][]int {
	inf := 1000000000
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = inf
		}
	}
	for _, road := range roads {
		a, b, c := road[0]-1, road[1]-1, road[2]
		graph[a][b] = min(graph[a][b], c)
		graph[b][a] = min(graph[b][a], c)
	}

	return graph
}

func min3(a, b int) int {
	if a < b {
		return a
	}
	return b
}
