package memo

func IsMatchRE(s string, p string) bool {
	memo := make([][]bool, len(s))    // 记忆搜索结果
	visited := make([][]bool, len(s)) // 标记是否访问
	for i := range memo {
		memo[i] = make([]bool, len(p))
		visited[i] = make([]bool, len(p))
	}

	return isMatchHelperRE(s, 0, p, 0, memo, visited)
}

func isMatchHelperRE(s string, sIndex int, p string, pIndex int, memo [][]bool, visited [][]bool) bool {
	// "" == ""
	if pIndex == len(p) { // 如果p已经匹配完毕
		return sIndex == len(s) // 根据s是否匹配完毕即可
	}

	if sIndex == len(s) { // 如果s匹配完毕
		return isEmpty(p, pIndex)
	}

	if visited[sIndex][pIndex] {
		return memo[sIndex][pIndex]
	}

	sChar := s[sIndex]
	pChar := p[pIndex]
	match := false

	// Consider a* as a bundle
	if pIndex+1 < len(p) && p[pIndex+1] == '*' { // 如果为'*'，有两种方案
		match = isMatchHelperRE(s, sIndex, p, pIndex+2, memo, visited) || // '*'不去匹配字符
			charMatchRE(sChar, pChar) && isMatchHelperRE(s, sIndex+1, p, pIndex, memo, visited) // '*'重复前面一个字符去匹配s
	} else {
		match = charMatchRE(sChar, pChar) && // 如果当前两字符匹配
			isMatchHelperRE(s, sIndex+1, p, pIndex+1, memo, visited) // 继续下一个字符匹配
	}

	visited[sIndex][pIndex] = true // 搜索完成就标记
	memo[sIndex][pIndex] = match   // 存储搜索结果
	return match
}

func charMatchRE(sChar, pChar byte) bool { // 判断两字符是否匹配
	return sChar == pChar || pChar == '.'
}

func isEmpty(p string, pIndex int) bool { // 形如"x*x*"形式
	for i := pIndex; i < len(p); i += 2 {
		if i+1 >= len(p) || p[i+1] != '*' { // 如果不是'*'，无法匹配
			return false
		}
	}
	return true
}
