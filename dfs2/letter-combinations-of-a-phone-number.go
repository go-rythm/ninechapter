package dfs2

// https://www.lintcode.com/problem/425/

/**
 * @param digits: A digital string
 * @return: all possible letter combinations
 *          we will sort your return value in output
 */

func LetterCombinations(digits string) []string {
	paths := []string{}
	if len(digits) == 0 {
		return paths
	}
	dict := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	dfs(dict, digits, 0, nil, &paths)
	return paths
}

// 递归三要素之一：递归的定义
// digits代表输入的数字
// idx代表当前dfs要处理digits[idx]所代表的数字
// path代表到目前为止得到的组合
// paths代表到目前为止找到的所有完整组合
func dfs(dict []string, digits string, idx int, path []rune, paths *[]string) {
	// 递归三要素之三：递归的出口
	// 已经找到了给定数字可以表示的一组combo，记录答案，并立即返回
	if idx == len(digits) {
		*paths = append(*paths, string(path))
		return
	}

	// 递归三要素之二：递归的拆解
	// 遍历idx位置的数字可以表示的不同字母
	for _, letter := range dict[digits[idx]-'2'] {
		path = append(path, letter)
		dfs(dict, digits, idx+1, path, paths)
		path = path[:len(path)-1]
	}
}
