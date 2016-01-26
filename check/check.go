/*
Package check - provides boolean tests to see if numbers meet criteria
such as primacy...
*/
package check

import (
	"math/big"
	"reflect"
)

// Prime returns true if n is prime
func Prime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(10)
}

// UniqueCharString returns true if a string contains no duplicate
// characters
func UniqueCharString(b []byte) bool {
	m := make(map[byte]int)
	for _, v := range b {
		m[v]++
		if m[v] > 1 {
			return false
		}
	}
	return true
}

// Palindrome returns true if b is a palindrome, false if not
func Palindrome(b []byte) bool {
	left, right := []byte{}, []byte{}
	if len(b)%2 == 0 {
		left = b[:len(b)/2]
		right = b[len(b)/2:]
	} else {
		left = b[:len(b)/2+1]
		right = b[len(b)/2:]
	}

	test := []byte{}
	for i := len(right) - 1; i >= 0; i-- {
		test = append(test, right[i])
	}

	if reflect.DeepEqual(left, test) {
		return true
	}
	return false
}
