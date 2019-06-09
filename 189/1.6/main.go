package main

import (
	"fmt"

	"algorithms/189/1.6/try1"
)

//String Compression: Implement a method to perform basic string compression using the counts
//of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
//"compressed" string would not become smaller than the original string, your method should return
//the original string. You can assume the string has only uppercase and lowercase letters (a - z).
func main() {
	str := []string{"aabcccccaaa", "AAADFDFGFFGGGYYYY", "", "aa", "b", "bcc", "cdef", "ffggg"}
	for _, s := range str {
		cs := try1.Compress(s)
		fmt.Printf("%s compressed to %s\n", s, cs)
	}
}
