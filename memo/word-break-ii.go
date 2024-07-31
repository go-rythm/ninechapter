package memo

func WordBreakII(s string, wordDict map[string]struct{}) []string {
	memo := make(map[string][]string)
	return dfs(s, wordDict, memo)
}

// 找到 s 的所有切割方案并 return
func dfs(s string, wordDict map[string]struct{}, memo map[string][]string) []string {
	if val, exists := memo[s]; exists {
		return val
	}

	if len(s) == 0 {
		return []string{}
	}

	partitions := []string{}

	for i := 1; i < len(s); i++ {
		prefix := s[:i]
		if _, ok := wordDict[prefix]; !ok {
			continue
		}

		subPartitions := dfs(s[i:], wordDict, memo)
		for _, partition := range subPartitions {
			partitions = append(partitions, prefix+" "+partition)
		}
	}

	if _, ok := wordDict[s]; ok {
		partitions = append(partitions, s)
	}

	memo[s] = partitions
	return partitions
}
