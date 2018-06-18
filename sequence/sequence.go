/*
Package sequence - functions for returning channels of sequences
*/
package sequence

import (
	"math"
	"math/big"

	"github.com/nboughton/numbers/isit"
)

// Primes returns a continuous channel of int64 Primes
func Primes() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		c <- 2

		for i := int64(3); i < int64(math.MaxInt64); i += 2 {
			if isit.Prime(i) {
				c <- i
			}
		}
	}()

	return c
}

// PrimesBetween returns a channel with all primes between start and finish
func PrimesBetween(start, finish int64) chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := range PrimeSieve(finish) {
			if i > start {
				c <- i
			}
		}
	}()

	return c
}

// PrimesFrom returns a channel of primes from start
func PrimesFrom(start int64) chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		if start == 2 {
			c <- start
		}

		if start%2 == 0 {
			start++
		}

		for i := start; i < int64(math.MaxInt64); i += 2 {
			if isit.Prime(i) {
				c <- i
			}
		}
	}()

	return c
}

// NPrimesFrom returns n conescutive primes starting from x
func NPrimesFrom(start, n int64) chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		if start == 2 {
			c <- start
		}

		if start%2 == 0 {
			start++
		}

		count := int64(0)
		for i := start; count < n; i += 2 {
			if isit.Prime(i) {
				c <- i
				count++
			}
		}
	}()

	return c
}

// PrimeSieve impements the Sieve of Eratosthenes and returns the streamed channel
// of results so that the program can cut off iteration where necessary
func PrimeSieve(value int64) chan int64 {
	f, c := make([]bool, value), make(chan int64)

	go func() {
		defer close(c)

		for i := int64(2); i <= int64(math.Sqrt(float64(value))); i++ {
			if f[i] == false {
				for j := i * i; j < value; j += i {
					f[j] = true
				}
			}
		}

		for i := int64(2); i < value; i++ {
			if !f[i] {
				c <- i
			}
		}
	}()

	return c
}

// PellLucas streams n iterations of the Pell/Pell-Lucas sequence. These can
// be used as approximations for the continued fraction of the square root of 2
func PellLucas(n int64) chan [2]*big.Int {
	a, b, r := big.NewInt(0), big.NewInt(1), make(chan [2]*big.Int)

	go func() {
		defer close(r)

		for i := int64(0); i < n; i++ {
			c, _ := big.NewInt(0).SetString(a.String(), 10)

			a.Add(a, big.NewInt(0).Mul(big.NewInt(2), b))
			a, b = b, a

			c.Add(c, a)

			r <- [2]*big.Int{big.NewInt(0).Set(c), big.NewInt(0).Set(a)}
		}
	}()

	return r
}

// Fibonacci returns a channel of the Fibonacci sequence using big Ints.
// Big ints are used because of the exponential growth of Fibonacci numbers.
func Fibonacci() chan *big.Int {
	c := make(chan *big.Int, 1)

	go func() {
		defer close(c)

		a, b := big.NewInt(0), big.NewInt(1)

		for true {
			a.Add(a, b)
			a, b = b, a
			c <- a
		}
	}()

	return c
}

// BigInts returns a continuous stream of big Ints integers from 1
func BigInts() chan *big.Int {
	c := make(chan *big.Int, 1)

	go func() {
		defer close(c)

		for i := big.NewInt(1); true; i.Add(i, big.NewInt(1)) {
			c <- i
		}
	}()

	return c
}

// Ints returns a continuous channel of integers from 1
func Ints() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i++ {
			c <- i
		}
	}()

	return c
}

// Evens returns a continuous channel of even numbers from 2
func Evens() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(2); i < int64(math.MaxInt64); i += 2 {
			c <- i
		}
	}()

	return c
}

// Odds returns a continuous channel of odd numbers from 1
func Odds() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i += 2 {
			c <- i
		}
	}()

	return c
}

// Triangles returns a channel of the triangle number sequence
func Triangles() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i++ {
			c <- i * (i + 1) / 2
		}
	}()

	return c
}

// Squares returns a channel of square numbers in sequence
func Squares() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i++ {
			c <- i * i
		}
	}()

	return c
}

// Pentagonals returns a channel of the pentagonal number sequence
func Pentagonals() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i++ {
			c <- i * (3*i - 1) / 2
		}
	}()

	return c
}

// Hexagonals returns a channel of the hexagonal number sequence
func Hexagonals() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i++ {
			c <- i * (2*i - 1)
		}
	}()

	return c
}

// Heptagonals returns a channel of the heptagonal number sequence
func Heptagonals() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i++ {
			c <- i * (5*i - 3) / 2
		}
	}()

	return c
}

// Octagonals returns a channel of the octagonal number sequence
func Octagonals() chan int64 {
	c := make(chan int64, 1)

	go func() {
		defer close(c)

		for i := int64(1); i < int64(math.MaxInt64); i++ {
			c <- i * (3*i - 2)
		}
	}()

	return c
}

// Rotations returns a sequence of rotations of n.
// I.e Rotations(123) = 123 -> 312 -> 231
func Rotations(n int64) chan int64 {
	rts := make(chan int64, 1)

	go func() {
		defer close(rts)

		s := []byte(big.NewInt(n).String())

		for i := 0; i < len(s); i++ {
			t := []byte{s[len(s)-1]}
			t = append(t, s[:len(s)-1]...)
			m, _ := big.NewInt(0).SetString(string(t), 10)
			rts <- m.Int64()
			s = t
		}
	}()

	return rts
}

// Truncate returns a channel of int64 slices that contain the
// truncation sequence of n from the left and the right simultaneously.
// I.e Truncate(123) = [123, 123] -> [23, 12] -> [3, 1]
func Truncate(n int64) chan []int64 {
	c := make(chan []int64, 1)

	go func() {
		d := []byte(big.NewInt(n).String())

		for i := range d {
			l, _ := big.NewInt(0).SetString(string(d[i:]), 10)
			r, _ := big.NewInt(0).SetString(string(d[:len(d)-i]), 10)
			c <- []int64{l.Int64(), r.Int64()}
		}

		close(c)
	}()

	return c
}

/* The functions below are channel/int64 adaptations of the Permutations and
Combinations functions found in github.com/ntns/goitertools/itertools */

// Permutations returns sucessive r length permutations of elements from
// iterable.
//
// Elements are treated as unique based on their position,
// not on their value. So if the input elements are unique, there
// will be no repeat values in each permutation.
//
//  Permutations([]int64{1, 2, 3}, 3) -> [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
func Permutations(iterable []int64, r int64) chan []int64 {
	ch := make(chan []int64, 1)
	if r > int64(len(iterable)) || r == 0 {
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)

		pool := iterable
		n := int64(len(pool))

		indices := make([]int64, n)
		for i := range indices {
			indices[i] = int64(i)
		}

		cycles := make([]int64, r)
		for i := range cycles {
			cycles[i] = n - int64(i)
		}

		result := make([]int64, r)
		for i, el := range indices[:r] {
			result[i] = pool[el]
		}
		ch <- result

		for n > 0 {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					index := indices[i]
					for j := i; j < n-1; j++ {
						indices[j] = indices[j+1]
					}
					indices[n-1] = index
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indices[i], indices[n-j] = indices[n-j], indices[i]

					result := make([]int64, r)
					for k := int64(0); k < r; k++ {
						result[k] = pool[indices[k]]
					}

					ch <- result
					break
				}
			}
			if i < 0 {
				break
			}
		}
	}()

	return ch
}

// Combinations returns r length subsquences of elements from
// iterable.
//
// Elements are treated as unique based on their position,
// not on their value. So if the input elements are unique, there
// will be no repeat values in each combination.
//  Combinations([]int64{1, 2, 3, 4, 5}, 4) -> [[1 2 3 4] [1 2 3 5] [1 2 4 5] [1 3 4 5] [2 3 4 5]]
func Combinations(iterable []int64, r int64) chan []int64 {
	ch := make(chan []int64, 1)
	if r > int64(len(iterable)) || r == 0 {
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)

		pool := iterable
		n := int64(len(pool))

		indices := make([]int64, r)
		for i := range indices {
			indices[i] = int64(i)
		}

		result := make([]int64, r)
		for i, el := range indices {
			result[i] = pool[el]
		}

		ch <- result

		for {
			i := r - 1
			for i >= 0 && indices[i] == i+n-r {
				i--
			}

			if i < 0 {
				break
			}

			indices[i]++
			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}

			result := make([]int64, r)
			for i = 0; i < int64(len(indices)); i++ {
				result[i] = pool[indices[i]]
			}

			ch <- result
		}
	}()

	return ch
}
