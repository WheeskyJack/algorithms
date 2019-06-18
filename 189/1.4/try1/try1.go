package try1

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
	unmatchedCount := 0
	var tmp rune
	length := len(s)
	for i, ch := range s {
		if i == 0 {
			tmp = ch
		}
		if ch != tmp {
			if (length - 1) == i { //end of for loop
				if i%2 != 0 { //even positioned character
					return false
				} else { //odd positioned character
					unmatchedCount++
					if unmatchedCount > 1 {
						return false
					} else {
						return true
					}
				}
			} else {
				if i%2 != 0 { //even positioned character
					unmatchedCount++
				} else {
					//do nothing
				}
			}
		} else { //character matched
			if (length - 1) == i { //end of for loop
				if i%2 != 0 { //even positioned character
					if unmatchedCount > 0 {
						return false
					} else {
						return true
					}
				} else { //odd positioned character
					if unmatchedCount > 1 {
						return false
					} else {
						return true
					}
				}
			} else { //not end of for loop
				if unmatchedCount > 1 { //fast return on failure
					return false
				} else {
					//do nothing
				}
			}
		}
		tmp = ch
	}
	return false
}

/*func isConsecutiveChar(s string) bool {
	unmatchedCount := 0
	var tmp rune
	length := len(s)
	for i, ch := range s {
		if i == 0 {
			tmp = ch
		}
		if i%2 != 0 { //even positioned character
			if ch != tmp {
				unmatchedCount++
			}
		} else {
			if (length - 1) == i { //final entry
				if ch != tmp {
					unmatchedCount++
				}
			}
		}
		if length%2 == 0 {
			if unmatchedCount > 0 {
				return false
			}
		} else {
			if unmatchedCount > 0 {
				return false
			}
		}
		tmp = ch
	}
	fmt.Println(unmatchedCount)
	return true
}
*/
