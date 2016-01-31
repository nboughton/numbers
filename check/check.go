/*
Package check - provides boolean tests to see if numbers meet criteria
such as primes, palindromes etc...
*/
package check

import (
	"github.com/nboughton/numbers/slice"
	"math/big"
	"reflect"
	"sort"
)

// Prime returns true if n is prime
func Prime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(10)
}

var panDigitalSets = map[int][]int64{
	1: {1},
	2: {1, 2},
	3: {1, 2, 3},
	4: {1, 2, 3, 4},
	5: {1, 2, 3, 4, 5},
	6: {1, 2, 3, 4, 5, 6},
	7: {1, 2, 3, 4, 5, 6, 7},
	8: {1, 2, 3, 4, 5, 6, 7, 8},
	9: {1, 2, 3, 4, 5, 6, 7, 8, 9},
}

// Pandigital returns true if n contains 1-n digits with each digit
// appearing only once
func Pandigital(n int) bool {
	t := slice.IntToSlice(n)
	sort.Ints(t)

	return reflect.DeepEqual(panDigitalSets[len(t)], t)
}

// PyTriplet returns true if a < b < c and a^2 + b^2 = c^2
func PyTriplet(a, b, c int64) bool {
	if a < b && b < c && a*a+b*b == c*c {
		return true
	}
	return false
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
	} else {
		left = b[:len(b)/2+1]
	}
	right = b[len(b)/2:]

	reverseRight := []byte{}
	for i := len(right) - 1; i >= 0; i-- {
		reverseRight = append(reverseRight, right[i])
	}

	return reflect.DeepEqual(left, reverseRight)
}
