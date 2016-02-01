/*
Package sequence - functions for returning channels of sequences
*/
package sequence

import (
	"github.com/nboughton/numbers/check"
	"github.com/nboughton/numbers/slice"
	"math/big"
)

// Primes returns a continuous channel of int64 Primes
func Primes() chan int64 {
	p := make(chan int64)

	go func() {
		p <- 2

		for i := int64(3); true; i += 2 {
			if check.Prime(i) {
				p <- i
			}
		}

		close(p)
	}()

	return p
}

// Fibonacci returns a channel of the Fibonacci sequence using big Ints
func Fibonacci() chan *big.Int {
	c := make(chan *big.Int)

	go func() {
		a, b := big.NewInt(0), big.NewInt(1)

		for true {
			a.Add(a, b)
			a, b = b, a
			c <- a
		}

		close(c)
	}()

	return c
}

// BigInts returns a continuous stream of big Ints integers from 1
func BigInts() chan *big.Int {
	c := make(chan *big.Int)
	i := big.NewInt(1)

	go func() {
		c <- i

		for true {
			i.Add(i, big.NewInt(1))
			c <- i
		}

		close(c)
	}()

	return c
}

// Ints returns a continuous channel of integers from 1
func Ints() chan int64 {
	c := make(chan int64)
	i := int64(1)

	go func() {
		c <- i

		for true {
			i++
			c <- i
		}

		close(c)
	}()

	return c
}

// Evens returns a continuous channel of even numbers from 2
func Evens() chan int64 {
	c := make(chan int64)
	i := int64(2)

	go func() {
		c <- i

		for true {
			i += 2
			c <- i
		}

		close(c)
	}()

	return c
}

// Odds returns a continuous channel of odd numbers from 1
func Odds() chan int64 {
	c := make(chan int64)
	i := int64(1)

	go func() {
		c <- i

		for true {
			i += 2
			c <- i
		}

		close(c)
	}()

	return c
}

// Triangles returns a channel of the triangle number sequence
func Triangles() chan int64 {
	c := make(chan int64)
	i := int64(1)

	go func() {
		for true {
			c <- i * (i + 1) / 2
			i++
		}
		close(c)
	}()

	return c
}

// Hexagonals returns a channel of the hexagonal number sequence
func Hexagonals() chan int64 {
	c := make(chan int64)
	i := int64(1)

	go func() {
		for true {
			c <- i * (2*i - 1)
			i++
		}
		close(c)
	}()

	return c
}

// Pentagonals returns a channel of the pentagonal number sequence
func Pentagonals() chan int64 {
	c := make(chan int64)
	i := int64(1)

	go func() {
		for true {
			c <- i * (3*i - 1) / 2
			i++
		}
		close(c)
	}()

	return c
}

// Rotations returns a sequence of rotations of n.
// I.e Rotations(123) = 123 -> 312 -> 231
func Rotations(n int64) chan []int64 {
	rts := make(chan []int64)

	go func() {
		a, b, c := slice.Int64ToSlice(n), []int64{}, []int64{}
		rts <- a

		for i := 1; i < len(a); i++ {
			if len(a) > 2 {
				b, c = a[len(a)-1:], a[:len(a)-1]
			} else if len(a) == 2 {
				b, c = a[1:], a[:1]
			}

			a = append(b, c...)
			rts <- a
		}

		close(rts)
	}()

	return rts
}

// Truncate returns a channel of int64 slices that contain the
// truncation sequence of n from the left and the right simultaneously.
// I.e Truncate(123) = [123, 123] -> [23, 12] -> [3, 1]
func Truncate(n int64) chan []int64 {
	c := make(chan []int64)

	go func() {
		d := slice.Int64ToSlice(n)

		for i := range d {
			c <- []int64{slice.SliceToInt64(d[i:]), slice.SliceToInt64(d[:len(d)-i])}
		}

		close(c)
	}()

	return c
}
