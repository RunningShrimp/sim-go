package utils

// FindCommonPrefix 寻找公共前缀
//
//	@Description:
//	@param s1
//	@param s2
//	@return int 相同字符数
func FindCommonPrefix(s1, s2 string) int {
	index := 0
	temp := s1
	strRune := []rune(s2)
	if len(s1) > len(s2) {
		temp = s2
		strRune = []rune(s1)
	}

	for i, ch := range temp {
		if strRune[i] != ch {
			return index
		}
		index = index + 1
	}
	return index

}
