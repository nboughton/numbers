/*
Package factor - provides functions for return factors, divisors and factorials
*/
package factor

import (
	//"fmt"
	"github.com/cznic/sortutil"
	"github.com/nboughton/numbers/check"
	"github.com/nboughton/numbers/total"
	"math"
	"math/big"
	"strconv"
)

// Primes returns int64 slice of prime factors of n
func Primes(n int64) []int64 {
	p1 := []int64{}
	t := int64(math.Sqrt(float64(n)))

	if n%2 == 0 {
		p1 = append(p1, int64(2))
	}

	if n < 1000 {
		t = n/2 + 1
	}

	for i := int64(3); i <= t; i += 2 {
		if n%i == 0 && check.Prime(i) {
			p1 = append(p1, i)
		}
	}

	return p1
}

// Divisors returns int64 slice of divisors of n
func Divisors(n int64) []int64 {
	f1, f2 := []int64{}, []int64{}
	t := int64(math.Sqrt(float64(n)))

	for i := int64(1); i <= t; i++ {
		if n%i == 0 {
			f1 = append(f1, i)
			if i*i != n {
				f2 = append(f2, n/i)
			}
		}
	}

	f1 = append(f1, f2...)
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
