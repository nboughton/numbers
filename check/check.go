/*
Package check - provides boolean tests to see if numbers meet criteria
such as primes, palindromes etc...
*/
package check

import (
	"math"
	"math/big"
	"reflect"
)

// Prime returns true if n is prime
func Prime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(10)
}

// PyTriplet returns true if a < b < c and a^2 + b^2 = c^2
func PyTriplet(a, b, c int64) bool {
	if a < b && b < c && a*a+b*b == c*c {
		return true
	}
	return false
}

// Triangular returns true is n is a triangular number
func Triangular(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(8*n+1)) - 1) / 2)
	if f == 0 && i > 1 {
		return true
	}
	return false
}

// Hexagonal returns true if n is a hexagonal number
func Hexagonal(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(8*n+1)) + 1) / 4)
	if f == 0 && i > 1 {
		return true
	}
	return false
}

// Pentagonal returns true if n is a pentagonal number
func Pentagonal(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(24*n+1)) + 1) / 6)
	if f == 0 && i > 0 {
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
