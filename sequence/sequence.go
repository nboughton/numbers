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

// Numbers returns a continuous channel of integers from 1
func Numbers() chan int64 {
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

// EvenNumbers returns a continuous channel of even numbers from 2
func EvenNumbers() chan int64 {
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

// OddNumbers returns a continuous channel of odd numbers from 1
func OddNumbers() chan int64 {
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

// Rotations returns a sequence of rotations of n
func Rotations(n int64) chan []int64 {
	rts := make(chan []int64)

	go func() {
		a, b, c := slice.IntToSlice(n), []int64{}, []int64{}
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
// truncation sequence of n from the left and the right simultaneously
func Truncate(n int64) chan []int64 {
	c := make(chan []int64)

	go func() {
		d := slice.IntToSlice(n)

		for i := range d {
			c <- []int64{slice.SliceToInt(d[i:]), slice.SliceToInt(d[:len(d)-i])}
		}

		close(c)
	}()

	return c
}
