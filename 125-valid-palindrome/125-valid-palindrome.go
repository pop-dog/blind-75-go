package main

import (
	"fmt"
	"unicode"

	"github.com/pop-dog/blind-75-go/dstructs/stack"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	san := ""
	// First pass: sanitize. convert all runes to lowercase and remove any non-alphanumeric
	for _, r := range s {
		if unicode.IsLetter(r) {
			san += string(unicode.ToLower(r))
		} else if unicode.IsNumber(r) {
			san += string(r)
		}
	}

	if len(san) == 0 {
		return true // null case. Empty string is palindrome
	}

	// Prep: determine midpoint
	// We either have
	// 				floor(n/2)	1							floor(n/2)
	// odd length: 	<left half>	<middle rune> (ignored>)	<right half> or,
	//				n/2			n/2
	// even length:	<left half> <right half>
	fhalf := san[:len(san)/2]
	var shalf string
	if len(san)%2 == 0 {
		shalf = san[len(san)/2:]
	} else {
		shalf = san[(len(san)/2)+1:]
	}

	// Second pass:
	// First half:	push to a stack
	rns := stack.New[rune]()
	for _, lr := range fhalf {
		rns.Push(lr)
	}
	// Second half:	pop from stack and compare
	for _, rr := range shalf {
		lr := rns.Pop()
		if lr != rr {
			return false
		}
	}
	return true
}

func main() {
	tests := []string{"A man, a plan, a canal: Panama", "race a car", " ", "0P"}
	for _, s := range tests {
		fmt.Println(s, ": ", isPalindrome(s))
	}
}
