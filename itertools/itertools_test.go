// Copyright 2012 Nuno Antunes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package itertools

import (
	"testing"
)

func TestCount(t *testing.T) {

	// should return nil

	if v := Count(1, 10, 0); v != nil {
		t.Errorf("Count(1, 10, 0) should return nil, got %v", v)
	}

	if v := Count(1, 10, -1); v != nil {
		t.Errorf("Count(1, 10, -1) should return nil, got %v", v)
	}

	if v := Count(10, 1, 1); v != nil {
		t.Errorf("Count(10, 1, 1) should return nil, got %v", v)
	}

	if v := Count(0, 0, 1); v != nil {
		t.Errorf("Count(1, 1, 1) should return nil, got %v", v)
	}

	// matches correct results

	if v := Count(1, 10, 1); !slice_match(v, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("Count(1, 10, 1) should return [1 2 3 4 5 6 7 8 9], got %v", v)
	}

	if v := Count(1, 10, 2); !slice_match(v, []int{1, 3, 5, 7, 9}) {
		t.Errorf("Count(1, 10, 2) should return [1 3 5 7 9], got %v", v)
	}

	if v := Count(10, 1, -1); !slice_match(v, []int{10, 9, 8, 7, 6, 5, 4, 3, 2}) {
		t.Errorf("Count(10, 1, -1) should return [10 9 8 7 6 5 4 3 2], got %v", v)
	}

	if v := Count(10, 1, -2); !slice_match(v, []int{10, 8, 6, 4, 2}) {
		t.Errorf("Count(10, 1, -2) should return [10 8 6 4 2], got %v", v)
	}

}

func TestCycle(t *testing.T) {

	// should return nil

	if v := Cycle([]int{1}, -1); v != nil {
		t.Errorf("Cycle([1], -1) should return nil, got %v", v)
	}

	if v := Cycle([]int{}, 1); v != nil {
		t.Errorf("Cycle([], 1) should return nil, got %v", v)
	}

	if v := Cycle(nil, 1); v != nil {
		t.Errorf("Cycle(nil, 1) should return nil, got %v", v)
	}

	// empty slice

	if v := Cycle([]int{1, 2, 3}, 0); !slice_match(v, []int{}) {
		t.Errorf("Cycle([1 2 3], 0) should return [], got %v", v)
	}

	// matches correct results

	if v := Cycle([]int{1, 2, 3, 4}, 6); !slice_match(v, []int{1, 2, 3, 4, 1, 2}) {
		t.Errorf("Cycle([1 2 3 4], 6) should return [1 2 3 4 1 2], got %v", v)
	}

	if v := Cycle([]int{1, 2, 3, 4, 5}, 3); !slice_match(v, []int{1, 2, 3}) {
		t.Errorf("Cycle([1 2 3 4 5], 3) should return [1 2 3], got %v", v)
	}

}

func TestRepeat(t *testing.T) {

	// should return nil

	if v := Repeat(10, -1); v != nil {
		t.Errorf("Repeat(10, -1) should return nil, got %v", v)
	}

	// empty slice

	if v := Repeat(10, 0); !slice_match(v, []int{}) {
		t.Errorf("Repeat(10, 0) should return [], got %v", v)
	}

	// matches correct results

	if v := Repeat(10, 10); !slice_match(v, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}) {
		t.Errorf("Repeat(10, 10) should return [10 10 10 10 10 10 10 10 10 10], got %v", v)
	}

}

func TestChain(t *testing.T) {

	// empty slice

	if v := Chain(); !slice_match(v, []int{}) {
		t.Errorf("Chain() should return [], got %v", v)
	}

	if v := Chain([]int{}); !slice_match(v, []int{}) {
		t.Errorf("Chain([]) should return [], got %v", v)
	}

	if v := Chain([]int{}, []int{}); !slice_match(v, []int{}) {
		t.Errorf("Chain([]) should return [], got %v", v)
	}

	// matches correct results

	if v := Chain([]int{1, 2, 3}); !slice_match(v, []int{1, 2, 3}) {
		t.Errorf("Chain([1 2 3]) should return [1 2 3], got %v", v)
	}

	if v := Chain([]int{1, 2, 3}, []int{}); !slice_match(v, []int{1, 2, 3}) {
		t.Errorf("Chain([1 2 3], []) should return [1 2 3], got %v", v)
	}

	if v := Chain([]int{1, 2, 3}, nil); !slice_match(v, []int{1, 2, 3}) {
		t.Errorf("Chain([1 2 3], nil) should return [1 2 3], got %v", v)
	}

	if v := Chain([]int{}, []int{4, 5, 6}); !slice_match(v, []int{4, 5, 6}) {
		t.Errorf("Chain([], [4 5 6]) should return [4 5 6], got %v", v)
	}

	if v := Chain(nil, []int{4, 5, 6}); !slice_match(v, []int{4, 5, 6}) {
		t.Errorf("Chain(nil, [4 5 6]) should return [4 5 6], got %v", v)
	}

	if v := Chain([]int{1, 2, 3}, []int{4, 5, 6}); !slice_match(v, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Chain([1 2 3], [4 5 6]) should return [1 2 3 4 5 6], got %v", v)
	}

	if v := Chain([]int{1, 2, 3}, []int{4, 5}, []int{6, 7, 8, 9}); !slice_match(v, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("Chain([1 2 3], [4 5] [6 7 8 9]) should return [1 2 3 4 5 6 7 8 9], got %v", v)
	}

	if v := Chain([]int{1, 2, 3}, []int{}, []int{4, 5, 6}); !slice_match(v, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Chain([1 2 3], [] [4 5 6]) should return [1 2 3 4 5 6], got %v", v)
	}

}

func TestCompress(t *testing.T) {

	// empty slice

	if v := Compress([]int{}, []int{}); !slice_match(v, []int{}) {
		t.Errorf("Compress([], []) should return [], got %v", v)
	}

	if v := Compress([]int{1, 2, 3}, []int{}); !slice_match(v, []int{}) {
		t.Errorf("Compress([1 2 3], []) should return [], got %v", v)
	}

	if v := Compress([]int{1, 2, 3}, nil); !slice_match(v, []int{}) {
		t.Errorf("Compress([1 2 3], nil) should return [], got %v", v)
	}

	if v := Compress([]int{}, []int{0, 1, 1}); !slice_match(v, []int{}) {
		t.Errorf("Compress([], [0 1 1]) should return [], got %v", v)
	}

	if v := Compress(nil, []int{0, 1, 1}); !slice_match(v, []int{}) {
		t.Errorf("Compress(nil, [0 1 1]) should return [], got %v", v)
	}

	// matches correct results

	if v := Compress([]int{1, 2, 3}, []int{0, 1, 1}); !slice_match(v, []int{2, 3}) {
		t.Errorf("Compress([1 2 3], [0 1 1]) should return [2 3], got %v", v)
	}

	if v := Compress([]int{1, 2, 3, 4, 5}, []int{1, 0, 1, 0}); !slice_match(v, []int{1, 3}) {
		t.Errorf("Compress([1 2 3 4 5], [1 0 1 0]) should return [1 3], got %v", v)
	}

	if v := Compress([]int{1, 2, 3}, []int{1, 0, 1, 1, 0, 1}); !slice_match(v, []int{1, 3}) {
		t.Errorf("Compress([1 2 3], [1 0 1 1 0 1]) should return [1 3], got %v", v)
	}

}

func TestDropWhile(t *testing.T) {

	// empty slice

	if v := DropWhile(is_odd, []int{}); !slice_match(v, []int{}) {
		t.Errorf("DropWhile(is_odd, []) should return [], got %v", v)
	}

	if v := DropWhile(is_odd, nil); !slice_match(v, []int{}) {
		t.Errorf("DropWhile(is_odd, nil) should return [], got %v", v)
	}

	if v := DropWhile(nil, []int{1, 2, 3}); !slice_match(v, []int{}) {
		t.Errorf("DropWhile(nil, [1 2 3]) should return [], got %v", v)
	}

	// matches correct results

	if v := DropWhile(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}); !slice_match(v, []int{2, 4, 5, 7, 6, 8}) {
		t.Errorf("DropWhile(is_odd, [1 3 2 4 5 7 6 8]) should return [2 4 5 7 6 8], got %v", v)
	}

}

func TestTakeWhile(t *testing.T) {

	// empty slice

	if v := TakeWhile(is_odd, []int{}); !slice_match(v, []int{}) {
		t.Errorf("TakeWhile(is_odd, []) should return [], got %v", v)
	}

	if v := TakeWhile(is_odd, nil); !slice_match(v, []int{}) {
		t.Errorf("TakeWhile(is_odd, nil) should return [], got %v", v)
	}

	if v := TakeWhile(nil, []int{1, 2, 3}); !slice_match(v, []int{}) {
		t.Errorf("TakeWhile(nil, [1 2 3]) should return [], got %v", v)
	}

	// matches correct results

	if v := TakeWhile(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}); !slice_match(v, []int{1, 3}) {
		t.Errorf("TakeWhile(is_odd, [1 3 2 4 5 7 6 8]) should return [1 3], got %v", v)
	}

}

func TestIFilter(t *testing.T) {

	// empty slice

	if v := IFilter(is_odd, []int{}); !slice_match(v, []int{}) {
		t.Errorf("IFilter(is_odd, []) should return [], got %v", v)
	}

	if v := IFilter(is_odd, nil); !slice_match(v, []int{}) {
		t.Errorf("IFilter(is_odd, nil) should return [], got %v", v)
	}

	if v := IFilter(nil, []int{-2, -1, 0, 1, 2}); !slice_match(v, []int{1, 2}) {
		t.Errorf("IFilter(nil, [-2 -1 0 1 2]) should return [1 2], got %v", v)
	}

	// matches correct results

	if v := IFilter(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}); !slice_match(v, []int{1, 3, 5, 7}) {
		t.Errorf("IFilter(is_odd, [1 3 2 4 5 7 6 8]) should return [1 3 5 7], got %v", v)
	}

}

func TestIFilterFalse(t *testing.T) {

	// empty slice

	if v := IFilterFalse(is_odd, []int{}); !slice_match(v, []int{}) {
		t.Errorf("IFilterFalse(is_odd, []) should return [], got %v", v)
	}

	if v := IFilterFalse(is_odd, nil); !slice_match(v, []int{}) {
		t.Errorf("IFilterFalse(is_odd, nil) should return [], got %v", v)
	}

	if v := IFilterFalse(nil, []int{-2, -1, 0, 1, 2}); !slice_match(v, []int{-2, -1, 0}) {
		t.Errorf("IFilterFalse(nil, [-2 -1 0 1 2]) should return [-2 -1 0], got %v", v)
	}

	// matches correct results

	if v := IFilterFalse(is_odd, []int{1, 3, 2, 4, 5, 7, 6, 8}); !slice_match(v, []int{2, 4, 6, 8}) {
		t.Errorf("IFilterFalse(is_odd, [1 3 2 4 5 7 6 8]) should return [2 4 6 8], got %v", v)
	}

}

func TestIZip(t *testing.T) {

	// should return nil

	if v := IZip(); v != nil {
		t.Errorf("IZip() should return nil, got %v", v)
	}

	// empty slice

	if v := IZip([]int{}); !slice2d_match(v, [][]int{}) {
		t.Errorf("IZip([]) should return [], got %v", v)
	}

	if v := IZip([]int{}, []int{}); !slice2d_match(v, [][]int{}) {
		t.Errorf("IZip([], []) should return [], got %v", v)
	}

	if v := IZip([]int{1, 2, 3}, []int{}); !slice2d_match(v, [][]int{}) {
		t.Errorf("IZip([1 2 3], []) should return [], got %v", v)
	}

	if v := IZip([]int{}, []int{1, 2, 3}); !slice2d_match(v, [][]int{}) {
		t.Errorf("IZip([], [1 2 3]) should return [], got %v", v)
	}

	// matches correct results

	if v := IZip([]int{1, 2}, []int{3, 4}); !slice2d_match(v, [][]int{{1, 3}, {2, 4}}) {
		t.Errorf("IZip([1 2], [3 4]) should return [[1 3] [2 4]], got %v", v)
	}

	if v := IZip([]int{1, 2}, []int{3, 4}, []int{5, 6}); !slice2d_match(v, [][]int{{1, 3, 5}, {2, 4, 6}}) {
		t.Errorf("IZip([1 2], [3 4] [5 6]) should return [[1 3 5] [2 4 6]], got %v", v)
	}

	if v := IZip([]int{1, 2, 3}, []int{4, 5}); !slice2d_match(v, [][]int{{1, 4}, {2, 5}}) {
		t.Errorf("IZip([1 2 3], [4 5]) should return [[1 4] [2 5]], got %v", v)
	}

	if v := IZip([]int{1, 2}, []int{3, 4, 5}); !slice2d_match(v, [][]int{{1, 3}, {2, 4}}) {
		t.Errorf("IZip([1 2], [3 4 5]) should return [[1 3] [2 4]], got %v", v)
	}

}

func TestIZipLongest(t *testing.T) {

	// should return nil

	if v := IZipLongest(-1); v != nil {
		t.Errorf("IZipLongest() should return nil, got %v", v)
	}

	// empty slice

	if v := IZipLongest(-1, []int{}); !slice2d_match(v, [][]int{}) {
		t.Errorf("IZipLongest([]) should return [], got %v", v)
	}

	if v := IZipLongest(-1, []int{}, []int{}); !slice2d_match(v, [][]int{}) {
		t.Errorf("IZipLongest([], []) should return [], got %v", v)
	}

	if v := IZipLongest(-1, []int{1, 2, 3}, []int{}); !slice2d_match(v, [][]int{{1, -1}, {2, -1}, {3, -1}}) {
		t.Errorf("IZipLongest([1 2 3], []) should return [[1 -1] [2 -1] [3 -1]], got %v", v)
	}

	if v := IZipLongest(-1, []int{}, []int{1, 2, 3}); !slice2d_match(v, [][]int{{-1, 1}, {-1, 2}, {-1, 3}}) {
		t.Errorf("IZip([], [1 2 3]) should return [[-1 1] [-1 2] [-1 3]], got %v", v)
	}

	// matches correct results

	if v := IZipLongest(-1, []int{1, 2}, []int{3, 4}); !slice2d_match(v, [][]int{{1, 3}, {2, 4}}) {
		t.Errorf("IZipLongest([1 2], [3 4]) should return [[1 3] [2 4]], got %v", v)
	}

	if v := IZipLongest(-1, []int{1, 2}, []int{3, 4}, []int{5, 6}); !slice2d_match(v, [][]int{{1, 3, 5}, {2, 4, 6}}) {
		t.Errorf("IZipLongest([1 2], [3 4] [5 6]) should return [[1 3 5] [2 4 6]], got %v", v)
	}

	if v := IZipLongest(-1, []int{1, 2}, []int{3, 4, 5}); !slice2d_match(v, [][]int{{1, 3}, {2, 4}, {-1, 5}}) {
		t.Errorf("IZipLongest([1 2], [3 4 5]) should return [[1 3] [2 4] [-1 5]], got %v", v)
	}

	if v := IZipLongest(-1, []int{1, 2, 3}, []int{4, 5}); !slice2d_match(v, [][]int{{1, 4}, {2, 5}, {3, -1}}) {
		t.Errorf("IZipLongest([1 2 3], [4 5]) should return [[1 4] [2 5] [3 -1]], got %v", v)
	}

}

func TestProduct(t *testing.T) {

	// empty product

	if v := Product(); !slice2d_match(v, [][]int{{}}) {
		t.Errorf("Product() should return [[]], got %v", v)
	}

	// should return nil

	if v := Product([]int{}); v != nil {
		t.Errorf("Product([]) should return nil, got %v", v)
	}

	if v := Product([]int{}, []int{}); v != nil {
		t.Errorf("Product([], []) should return nil, got %v", v)
	}

	if v := Product([]int{}, []int{1, 2, 3}); v != nil {
		t.Errorf("Product([], [1 2 3]) should return nil, got %v", v)
	}

	if v := Product([]int{1, 2, 3}, []int{}); v != nil {
		t.Errorf("Product([1 2 3], []) should return nil, got %v", v)
	}

	// matches correct results

	if v := Product([]int{1, 2, 3}); !slice2d_match(v, [][]int{{1}, {2}, {3}}) {
		t.Errorf("Product([1 2 3]) should return [[1] [2] [3]], got %v", v)
	}

	if v := Product([]int{1, 2}, []int{3, 4}); !slice2d_match(v, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}}) {
		t.Errorf("Product([1 2], [3 4]) should return [[1 3] [1 4] [2 3] [2 4]], got %v", v)
	}

	// length of results is correct

	if v := len(Product([]int{1, 2, 3}, []int{4, 5, 6})); v != 9 {
		t.Errorf("len(Product([1 3 4], [4 5 6])) should return 9, got %v", v)
	}

	if v := len(Product([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8})); v != 18 {
		t.Errorf("len(Product([1 3 3], [4 5 6], [7 8])) should return 18, got %v", v)
	}

}

func TestPermutations(t *testing.T) {

	// should return nil

	if v := Permutations([]int{}, 0); v != nil {
		t.Errorf("Permutations([], 0) should return nil, got %v", v)
	}

	if v := Permutations([]int{}, 5); v != nil {
		t.Errorf("Permutations([], 5) should return nil, got %v", v)
	}

	if v := Permutations([]int{1, 2, 3}, 5); v != nil {
		t.Errorf("Permutations([1 2 3], 5) should return nil, got %v", v)
	}

	// matches correct results

	if v := Permutations([]int{1, 2, 3}, 3); !slice2d_match(v, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}) {
		t.Errorf("Permutations([1 2 3], 3) should return [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]], got %v", v)
	}

	// length of results is correct

	if v := len(Permutations([]int{1, 2, 3, 4}, 3)); v != 24 {
		t.Errorf("len(Permutations([1 2 3 4], 3)) should return 20, got %v", v)
	}

	if v := len(Permutations([]int{1, 2, 3, 4, 5, 6}, 5)); v != 720 {
		t.Errorf("len(Permutations([1 2 3 4 5 6], 5)) should return 252, got %v", v)
	}

}

func TestCombinations(t *testing.T) {

	// should return nil

	if v := Combinations([]int{}, 0); v != nil {
		t.Errorf("Combinations([], 0) should return nil, got %v", v)
	}

	if v := Combinations([]int{}, 5); v != nil {
		t.Errorf("Combinations([], 5) should return nil, got %v", v)
	}

	if v := Combinations([]int{1, 2, 3}, 5); v != nil {
		t.Errorf("Combinations ([1 2 3], 5) should return nil, got %v", v)
	}

	// matches correct results

	if v := Combinations([]int{1, 2, 3}, 3); !slice2d_match(v, [][]int{{1, 2, 3}}) {
		t.Errorf("Combinations([1 2 3], 3) should return [[1 2 3]], got %v", v)
	}

	if v := Combinations([]int{1, 2, 3, 4, 5}, 4); !slice2d_match(v, [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}}) {
		t.Errorf("Combinations([1 2 3 4 5], 4) should return [[1 2 3 4] [1 2 3 5] [1 2 4 5] [1 3 4 5] [2 3 4 5]], got %v", v)
	}

	// length of results is correct

	if v := len(Combinations([]int{1, 2, 3, 4, 5, 6}, 3)); v != 20 {
		t.Errorf("len(Combinations([1 2 3 4 5 6], 3)) should return 20, got %v", v)
	}

	if v := len(Combinations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)); v != 252 {
		t.Errorf("len(Combinations([1 2 3 4 5 6 7 8 9 10], 5)) should return 252, got %v", v)
	}

}

// helper functions

func slice_match(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func slice2d_match(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, s := range a {
		if len(s) != len(b[i]) {
			return false
		}

		for j, v := range s {
			if v != b[i][j] {
				return false
			}
		}
	}

	return true
}

func is_odd(v int) bool {

	if v%2 != 0 {
		return true
	}

	return false
}
