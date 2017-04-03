/*
Package slice - functions for converting numbers to and from slices
*/
package slice

import (
	"math/big"
	"strconv"
)

// ToBigInt returns a big Int composite of the digits in n
func ToBigInt(n []int64) *big.Int {
	s := []byte{}

	for _, v := range n {
		s = strconv.AppendInt(s, v, 10)
	}

	i, _ := big.NewInt(0).SetString(string(s), 10)
	return i
}

// ToInt64 takes a slice of int64s and returns their composite int64
func ToInt64(n []int64) int64 {
	s := []byte{}

	for _, v := range n {
		s = strconv.AppendInt(s, v, 10)
	}

	i, _ := strconv.ParseInt(string(s), 10, 64)
	return i
}

// ToInt takes a slice of ints and returns their composite int
func ToInt(n []int) int {
	s := []byte{}

	for _, v := range n {
		s = strconv.AppendInt(s, int64(v), 10)
	}

	i, _ := strconv.ParseInt(string(s), 10, 32)
	return int(i)
}

// FromInt64 returns a number as a slice of its digits
func FromInt64(n int64) []int64 {
	nInts := []int64{}

	for _, v := range strconv.FormatInt(n, 10) {
		i, _ := strconv.ParseInt(string(v), 10, 64)
		nInts = append(nInts, i)
	}

	return nInts
}

// FromInt returns a number as a slice of its digits
func FromInt(n int) []int {
	nInts := []int{}

	for _, v := range strconv.Itoa(n) {
		i, _ := strconv.Atoi(string(v))
		nInts = append(nInts, i)
	}

	return nInts
}
