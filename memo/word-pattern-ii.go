package memo

import (
	"strings"
)

func WordPatternMatch(pattern string, str string) bool {
	return isMatch(pattern, str, map[rune]string{}, map[string]struct{}{})
}

func isMatch(pattern string, str string, mapping map[rune]string, used map[string]struct{}) bool {
	if len(pattern) == 0 {
		return len(str) == 0
	}

	char := rune(pattern[0])
	if word, ok := mapping[char]; ok {
		if !strings.HasPrefix(str, word) {
			return false
		}
		return isMatch(pattern[1:], str[len(word):], mapping, used)
	}

	for i := 0; i < len(str); i++ {
		word := str[:i+1]
		if _, exists := used[word]; exists {
			continue
		}

		used[word] = struct{}{}
		mapping[char] = word

		if isMatch(pattern[1:], str[i+1:], mapping, used) {
			return true
		}

		delete(mapping, char)
		delete(used, word)
	}

	return false
}
