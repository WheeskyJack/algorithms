package main

import (
	"fmt"

	"algorithms/189/1.2/try1"
)

//Check Permutation: Given two strings, write a method to decide if one is a permutation of the
//other.
func main() {
	str1 := []string{"abc", "defg", "hijk"}
	str2 := []string{"bca", "dgef", "jkhi"}
	for _, s1 := range str1 {
		for _, s2 := range str2 {
			if isMatch := try1.IsPermutation(s1, s2); isMatch {
				fmt.Printf("%s macthes %s\n", s1, s2)
			} else {
				fmt.Printf("%s doesn't macth %s\n", s1, s2)
			}
		}
	}
}
