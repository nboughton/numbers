/*
Package slice - functions for converting numbers to and from slices
*/
package slice

import (
	"fmt"
	"strconv"
	"strings"
)

// SliceToInt takes a slice of ints and returns their composite integer
func SliceToInt(n []int64) int64 {
	s := ""

	for _, v := range n {
		s = s + fmt.Sprintf("%v", v)
	}

	i, _ := strconv.Atoi(s)
	return int64(i)
}

// IntToSlice returns a number as a slice of its digits
func IntToSlice(n int64) []int64 {
	nStrs := strings.Split(fmt.Sprintf("%v", n), "")
	nInts := []int64{}

	for _, v := range nStrs {
		i, _ := strconv.Atoi(v)
		nInts = append(nInts, int64(i))
	}

	return nInts
}
