package try2

import (
	"sort"
	"strings"
	"unicode"
)

func CanFormPalindrome(s string) bool {
	s = sortString(StripSpaceFromString(s))
	return isConsecutiveChar(s)
}

func StripSpaceFromString(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return strings.ToLower(b.String())
}

func sortString(s string) string {
	sA := strings.Split(s, "")
	sort.Strings(sA)
	return strings.Join(sA, "")
}

func isConsecutiveChar(s string) bool {
	m := make(map[rune]int, 0)
	for _, ch := range s {
		if _, ok := m[ch]; ok {
			delete(m, ch)
		} else {
			m[ch] = 1
		}
	}
	if len(s)%2 == 0 { //even lenght string
		if len(m) == 0 {
			return true
		} else {
			return false
		}
	} else {
		if len(m) == 1 {
			return true
		} else {
			return false
		}
	}
}
