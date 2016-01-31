/*
Package slice - functions for converting numbers to and from slices
*/
package slice

import (
	"math/big"
	"strconv"
	"strings"
)

// SliceToBigInt returns a big Int composite of the digits in n
func SliceToBigInt(n []int) *big.Int {
	s := ""

	for _, v := range n {
		s += strconv.Itoa(v)
	}

	i, _ := big.NewInt(0).SetString(s, 10)
	return i
}

// SliceToInt64 takes a slice of int64s and returns their composite int64
func SliceToInt64(n []int64) int64 {
	s := ""

	for _, v := range n {
		s += strconv.FormatInt(v, 10)
	}

	i, _ := strconv.Atoi(s)
	return int64(i)
}

// SliceToInt64 takes a slice of ints and returns their composite int
func SliceToInt(n []int) int {
	s := ""

	for _, v := range n {
		s += strconv.Itoa(v)
	}

	i, _ := strconv.Atoi(s)
	return i
}

// Int64ToSlice returns a number as a slice of its digits
func Int64ToSlice(n int64) []int64 {
	nStrs := strings.Split(strconv.FormatInt(n, 10), "")
	nInts := []int64{}

	for _, v := range nStrs {
		i, _ := strconv.Atoi(v)
		nInts = append(nInts, int64(i))
	}

	return nInts
}

// IntToSlice returns a number as a slice of its digits
func IntToSlice(n int) []int {
	nStrs := strings.Split(strconv.Itoa(n), "")
	nInts := []int{}

	for _, v := range nStrs {
		i, _ := strconv.Atoi(v)
		nInts = append(nInts, i)
	}

	return nInts
}
