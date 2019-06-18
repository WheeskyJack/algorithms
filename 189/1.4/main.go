package main

import (
	"fmt"

	"algorithms/189/1.4/try2"
)

//Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palin­
//drome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation
//is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
func main() {
	str1 := []string{"TactCoa", "Tact Coa",
		"abc", "pqrs", "b  c", "abac",
		"aaa", "b", "b   b", "aba", "dd", "abacc",
		"abb", "bba", "aabbc"}
	for _, s := range str1 {
		if ok := try2.CanFormPalindrome(s); ok {
			fmt.Printf("%s is a permutation of a Palindrome\n", s)
		} else {
			fmt.Printf("%s is not a permutation of a Palindrome\n", s)
		}
	}
}
