package try1

import (
	"strings"
)

func OneWay(s1, s2 string) bool {
	ld := len(s2) - len(s1)
	if ld > 1 { //fast return
		return false
	} else if ld > 0 {
		s1, s2 = s2, s1
	} else {
		if s1 == s2 {
			return true
		}
		//do nothing
	}
	s1A := strings.Split(s1, "")
	s2A := strings.Split(s2, "")
	var s2UnmatchIndex = len(s2A)
	var s1Char string

	for ind, s2ch := range s2A {
		if s2ch != s1A[ind] {
			s2UnmatchIndex = ind
			s1Char = s1A[ind]
			break
		} else {
			//do nothing
		}
	}

	s1del := getDelString(s1A, s2UnmatchIndex)
	if s1del == s2 {
		return true
	}
	s2Rpl := getReplaceString(s2A, s2UnmatchIndex, s1Char)
	if s2Rpl == s1 {
		return true
	}
	return false
}

func getDelString(s []string, ind int) string {
	var ret string
	for i, sc := range s {
		if i == ind {
			continue
		}
		ret += sc
	}
	return ret
}

func getReplaceString(s []string, ind int, sc string) string {
	s[ind] = sc
	return strings.Join(s, "")
}
