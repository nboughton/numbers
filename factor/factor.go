/*
Package factor - provides functions for return factors, divisors and factorials
*/
package factor

import (
	"math"
	"math/big"
	"strconv"

	"github.com/nboughton/numbers/isit"
	"github.com/nboughton/numbers/set"
	"github.com/nboughton/numbers/total"
)

// Primes returns int64 slice of prime factors of n
func Primes(n int64) set.Int64 {
	p := set.Int64{}

	for _, v := range Divisors(n) {
		if isit.Prime(v) {
			p = append(p, v)
		}
	}

	return p
}

// Divisors returns int64 slice of divisors of n
func Divisors(n int64) set.Int64 {
	f, t := set.Int64{}, int64(math.Sqrt(float64(n)))

	for i := int64(1); i <= t; i++ {
		if n%i == 0 {
			f = append(f, i)
			if i*i != n {
				f = append(f, n/i)
			}
		}
	}

	return f.Dedupe()
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

// Totient returns the result of Eulers Totient or Phi function of value n
func Totient(n int64) int64 {
	pF := Primes(n)

	ans := n

	for _, prime := range pF {
		ans = ans * (prime - 1) / prime
	}

	return ans
}
