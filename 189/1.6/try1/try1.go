package try1

import (
	"strconv"
)

func Compress(s string) string {
	if len(s) == 0 {
		return s
	}
	pc := rune(s[0])
	count := 0
	r := ""
	for _, c := range s {
		if pc != c {
			r += string(pc)
			r += strconv.Itoa(count)
			count = 0
		}
		count++
		pc = c
	}
	r += string(pc)
	r += strconv.Itoa(count)
	if len(r) >= len(s) {
		return s
	}
	return r
}
