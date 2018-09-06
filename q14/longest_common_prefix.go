package q14

// 14. Longest Common Prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	baseStr := strs[0]
	for _, str := range strs {
		if len(str) < len(baseStr) {
			baseStr = str
		}
	}
	for i := 0; i < len(baseStr); i++ {
		for j := 0; j < len(strs); j++ {
			if baseStr[i] != strs[j][i] {
				return res
			}
		}
		res += string(baseStr[i])
	}
	return res
}
