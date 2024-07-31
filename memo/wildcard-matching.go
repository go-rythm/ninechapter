package memo

/**
 * @param s: A string
 * @param p: A string includes "?" and "*"
 * @return: is Match?
 */
func IsMatch(s string, p string) bool {
	memo := make([][]bool, len(s))
	visited := make([][]bool, len(s))
	for i := range memo {
		memo[i] = make([]bool, len(p))
		visited[i] = make([]bool, len(p))
	}

	return isMatchHelper(s, 0, p, 0, memo, visited)
}

func isMatchHelper(s string, sIndex int, p string, pIndex int, memo [][]bool, visited [][]bool) bool {
	// 如果 p 从pIdex开始是空字符串了，那么 s 也必须从 sIndex 是空才能匹配上
	if pIndex == len(p) {
		return sIndex == len(s)
	}

	// 如果 s 从 sIndex 是空，那么p 必须全是 *
	if sIndex == len(s) {
		return allStar(p, pIndex)
	}

	if visited[sIndex][pIndex] {
		return memo[sIndex][pIndex]
	}

	sChar := s[sIndex]
	pChar := p[pIndex]
	match := false

	if pChar == '*' {
		match = isMatchHelper(s, sIndex, p, pIndex+1, memo, visited) ||
			isMatchHelper(s, sIndex+1, p, pIndex, memo, visited)
	} else {
		match = charMatch(sChar, pChar) &&
			isMatchHelper(s, sIndex+1, p, pIndex+1, memo, visited)
	}

	visited[sIndex][pIndex] = true
	memo[sIndex][pIndex] = match
	return match
}

func charMatch(sChar, pChar byte) bool {
	return sChar == pChar || pChar == '?'
}

func allStar(p string, pIndex int) bool {
	for i := pIndex; i < len(p); i++ {
		if p[i] != '*' {
			return false
		}
	}
	return true
}
