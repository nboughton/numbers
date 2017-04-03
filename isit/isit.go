/*
Package isit - provides boolean tests to see if numbers meet criteria
such as primes, palindromes etc...
*/
package isit

import (
	"math"
	"math/big"
)

// Prime returns true if n is prime
func Prime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(5)
}

// PyTriplet returns true if a < b < c and a^2 + b^2 = c^2
func PyTriplet(a, b, c int64) bool {
	if a < b && b < c && a*a+b*b == c*c {
		return true
	}
	return false
}

// Triangular returns true if n is a triangular number
func Triangular(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(8*n+1)) - 1) / 2)
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// Square returns true if n is a square number
func Square(n int64) bool {
	i, f := math.Modf(math.Sqrt(float64(n)))
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// Pentagonal returns true if n is a pentagonal number
func Pentagonal(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(24*n+1)) + 1) / 6)
	if f == 0 && i >= 0 {
		return true
	}
	return false
}

// Hexagonal returns true if n is a hexagonal number
func Hexagonal(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(8*n+1)) + 1) / 4)
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// Heptagonal returns true if n is a heptagonal number
func Heptagonal(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(40*n+9)) + 3) / 10)
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// Octagonal returns true if n is an octagonal number
func Octagonal(n int64) bool {
	i, f := math.Modf((math.Sqrt(float64(3*n+1)) + 1) / 3)
	if f == 0 && i >= 1 {
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

// Palindrome returns true if b is a palindrome
func Palindrome(b []byte) bool {
	dst := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		dst[i] = b[len(b)-1-i]
	}

	return string(dst) == string(b)
}
