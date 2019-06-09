package try1

import (
	"sort"
	"strings"
)

func IsPermutation(s1, s2 string) bool {
	s1 = sortString(s1)
	s2 = sortString(s2)
	if s1 == s2 {
		return true
	}
	return false
}

func sortString(s string) string {
	sA := strings.Split(s, "")
	sort.Strings(sA)
	return strings.Join(sA, "")
}
