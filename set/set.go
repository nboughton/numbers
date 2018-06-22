package set

import (
	"sort"
)

// Int64 is a slice of int64
type Int64 []int64

func (s Int64) Len() int           { return len(s) }
func (s Int64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int64) Less(i, j int) bool { return s[i] < s[j] }

// Contains returns whether or not n exists in set s
func (s Int64) Contains(n int64) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}

	return false
}

// Sum returns the sum total of the set
func (s Int64) Sum() int64 {
	var t int64

	for _, n := range s {
		t += n
	}

	return t
}

// Dedupe returns a set with only unique values
func (s Int64) Dedupe() Int64 {
	var (
		m   = make(map[int64]int)
		res Int64
	)

	for _, n := range s {
		m[n]++
	}

	for k := range m {
		res = append(res, k)
	}

	sort.Sort(res)

	return res
}

// Int64s is a slice of slices of int64
type Int64s []Int64

func (s Int64s) Len() int           { return len(s) }
func (s Int64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int64s) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

// Find returns all sets in s that contain n
func (s Int64s) Find(n int64) chan Int64 {
	c := make(chan Int64)

	go func() {
		defer close(c)

		for _, set := range s {
			if set.Contains(n) {
				c <- set
			}
		}
	}()

	return c
}
