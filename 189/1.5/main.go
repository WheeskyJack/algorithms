package main

import (
	"fmt"

	"algorithms/189/1.5/try1"
)

//One Away: There are three types of edits that can be performed on strings: insert a character,
//remove a character, or replace a character. Given two strings, write a function to check if they are
//one edit (or zero edits) away.
func main() {
	str := [][]string{
		{"pale", "ple"},
		{"pale", "ale"},
		{"pale", "pal"},
		{"ple", "pale"},
		{"ple", "ple"},
		{"pales", "pale"},
		{"pale", "bale"},
		{"pale", "bake"},
		{"pale", "ake"},
		{"pale", "al"},
	}
	for _, s := range str {
		if !try1.OneWay(s[0], s[1]) {
			fmt.Printf("%s %s are not one edit away\n", s[0], s[1])
		} else {
			fmt.Printf("%s %s are one edit away\n", s[0], s[1])
		}
	}
}
