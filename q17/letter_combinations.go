package q17

var phoneLetterMap = map[uint8][]uint8{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) <= 0 {
		return res
	}
	res = append(res, "")
	for len(res[0]) != len(digits) {
		peek := res[0]
		letters := phoneLetterMap[digits[len(peek)]]
		res = res[1:]
		for _, c := range letters {
			res = append(res, peek+string(c))
		}
	}
	return res
}
