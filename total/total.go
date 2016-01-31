/*
Package total - Provides functions for sums, products, powers etc
*/
package total

import (
	//"math"
	"math/big"
	"strings"
)

var abc = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var atoi = make(map[byte]int)

func init() {
	for i, v := range abc {
		atoi[v] = i + 1
	}
}

// Abc returns the a = 1, b = 2 etc score for string s
func Abc(s string) int64 {
	s = strings.ToUpper(s)

	score := 0
	for _, v := range []byte(s) {
		score += atoi[v]
	}

	return int64(score)
}

// Sum returns sum of terms n
func Sum(n []int64) int64 {
	s := int64(0)
	for _, v := range n {
		s += v
	}
	return s
}

// BigSum takes an array of strings representing big ints and returns
// a big Int value of the sum
func BigSum(n []string) *big.Int {
	a := big.NewInt(0)
	a.SetString(n[0], 10)

	for i := 1; i < len(n); i++ {
		b := big.NewInt(0)
		b.SetString(n[i], 10)

		a.Add(a, b)
	}
	return a
}

// Product returns product of terms n
func Product(n []int64) int64 {
	p := int64(n[0])
	for i := 1; i < len(n); i++ {
		p *= n[i]
	}
	return p
}

// BigProduct takes an array of strings representing numbers and
// and returns a big Int containing their Product
func BigProduct(n []string) *big.Int {
	a := big.NewInt(0)
	a.SetString(n[0], 10)

	for i := 1; i < len(n); i++ {
		b := big.NewInt(0)
		b.SetString(n[i], 10)

		a.Mul(a, b)
	}
	return a
}

// BigPow returns x^y as a big Int
func BigPow(x, y int64) *big.Int {
	n, m := big.NewInt(x), big.NewInt(x)
	for i := int64(2); i <= y; i++ {
		n.Mul(n, m)
	}
	return n
}
