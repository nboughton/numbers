/*
Package factor - provides functions for return factors, divisors and factorials
*/
package factor

import (
	//"fmt"
	"github.com/cznic/sortutil"
	"github.com/nboughton/numbers/isit"
	"github.com/nboughton/numbers/total"
	"math"
	"math/big"
	"strconv"
)

// Primes returns int64 slice of prime factors of n
func Primes(n int64) []int64 {
	p := []int64{}

	for _, v := range Divisors(n) {
		if isit.Prime(v) {
			p = append(p, v)
		}
	}

	return p
}

// Divisors returns int64 slice of divisors of n
func Divisors(n int64) []int64 {
	f1 := []int64{}
	t := int64(math.Sqrt(float64(n)))

	for i := int64(1); i <= t; i++ {
		if n%i == 0 {
			f1 = append(f1, i)
			if i*i != n {
				f1 = append(f1, n/i)
			}
		}
	}

	sortutil.Int64Slice(f1).Sort()
	sortutil.Dedupe(sortutil.Int64Slice(f1))
	return f1
}

// Factorial returns n! using big Ints
func Factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	set := []string{}
	for i := n; i > 0; i-- {
		set = append(set, strconv.FormatInt(i, 10))
	}
	return total.BigProduct(set)
}
