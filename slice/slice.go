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
	s := ""

	for _, v := range n {
		s += strconv.FormatInt(v, 10)
	}

	i, _ := big.NewInt(0).SetString(s, 10)
	return i
}

// ToInt64 takes a slice of int64s and returns their composite int64
func ToInt64(n []int64) int64 {
	s := ""

	for _, v := range n {
		s += strconv.FormatInt(v, 10)
	}

	i, _ := strconv.Atoi(s)
	return int64(i)
}

// ToInt takes a slice of ints and returns their composite int
func ToInt(n []int) int {
	s := ""

	for _, v := range n {
		s += strconv.Itoa(v)
	}

	i, _ := strconv.Atoi(s)
	return i
}

// FromInt64 returns a number as a slice of its digits
func FromInt64(n int64) []int64 {
	nStrs, nInts := []byte((strconv.FormatInt(n, 10))), []int64{}

	for _, v := range nStrs {
		i, _ := strconv.Atoi(string(v))
		nInts = append(nInts, int64(i))
	}

	return nInts
}

// FromInt returns a number as a slice of its digits
func FromInt(n int) []int {
	nStrs, nInts := []byte((strconv.Itoa(n))), []int{}

	for _, v := range nStrs {
		i, _ := strconv.Atoi(string(v))
		nInts = append(nInts, i)
	}

	return nInts
}
