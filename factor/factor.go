/*
Package factor - provides functions for return factors, divisors and factorials
*/
package factor

import (
	"github.com/nboughton/numbers/check"
	"github.com/nboughton/numbers/total"
	"math"
	"math/big"
	"sort"
	"strconv"
)

// Primes returns int64 slice of prime factors of n
func Primes(n int64) []int64 {
	p := []int64{}
	for i := int64(3); float64(i) < math.Ceil(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 && check.Prime(i) {
			p = append(p, i)
		}
	}
	return p
}

// Divisors returns int64 slice of divisors of n
func Divisors(n int64) []int64 {
	f := []int64{}
	lf := []int{}
	for i := int64(1); float64(i) <= math.Ceil(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			f = append(f, i)
			if i*i != n {
				lf = append(lf, int(n/i))
			}
		}
	}
	sort.Ints(lf)
	for _, v := range lf {
		dup := false
		for _, r := range f {
			if int64(v) == r {
				dup = true
				break
			}
		}
		if !dup {
			f = append(f, int64(v))
		}
	}
	return f
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
