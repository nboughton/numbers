package set

import (
	"fmt"
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

// Dedupe returns a sorted set with only unique values
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

// MaxPathSum returns the maximum value available in a path through
// a numerical grid, i.e a Set of sets
func (s Int64s) MaxPathSum() int64 {
	for row := len(s) - 2; row >= 0; row-- {
		for col := 0; col < len(s[row])-1; col++ {
			if s[row+1][col] > s[row+1][col+1] {
				s[row][col] += s[row+1][col]
			} else {
				s[row][col] += s[row+1][col+1]
			}
		}
	}

	return s[0][0]
}

/* Using Int64s as a grid
Consider the following:
[]int{
	[]int{25,10,11,12,13},
	[]int{24,09,02,03,14},
	[]int{23,08,01,04,15},
	[]int{22,07,06,05,16},
	[]int{21,20,19,18,17}
}
*/

// NewNumberSpiral creates a square grid number spiral of width size. If size is even it is incremented
// to become odd.
/*
CONSIDER:
for i := 1; i < max; i += inc {
	inc increases every 2nd and 4th turn
	use vector, supply n = i..{i+inc}
}

func NewNumberSpiral(size int64) Int64s {
	if size%2 == 0 {
		size++
	}

	grid := make(Int64s, size)
	for row := range grid {
		grid[row] = make(Int64, size)
	}

	// Starting from the center head up 1...
	row, col, max := size/2, size/2, size*size
	inc, turn := int64(1), int64(0)

	for i := int64(1); i <= max; i++ {

		// Print update after each line
		for _, r := range grid {
			fmt.Println(r)
		}
	}

	return grid
}
*/
/*
	switch {
			case d == UP:
				grid[row-1][col] = i
			case d == LTR:
				grid[row][col+1] = i
			case d == DOWN:
				grid[row+1][col] = i
			case d == RTL:
				grid[row][col-1] = i
			}
*/

// Direction represents an identifier for vector direction
type Direction int

// Coord represents the values of coordinates within a grid
type Coord struct {
	Row int64
	Col int64
}

// Vector Directions constants
const (
	LTR  Direction = iota // Left To Right
	RTL                   // Right To Left
	UP                    // Up
	DOWN                  // Down
	LTRU                  // Left To Right Up (diagonal)
	LTRD                  // Left To Right Down (diagonal)
	RTLU                  // Right To Left Up (diagonal)
	RTLD                  // Right To Left Down (diagonal)
)

// Vector returns a ln length set of values starting at row/col extending in Direction d.
// Vector also returns the coordinates of those values.
// If supplied Vector will set the values to n (in order)
func (s Int64s) Vector(r, c, ln int64, d Direction, n ...int64) (Int64, []Coord, error) {
	var (
		res  Int64
		crds = make([]Coord, ln)
	)

	for i := int64(0); i < ln; i++ {
		crd := Coord{}

		switch d {
		case LTR:
			crd = Coord{r, c + i}
		case RTL:
			crd = Coord{r, c - i}
		case DOWN:
			crd = Coord{r + i, c}
		case UP:
			crd = Coord{r - i, c}
		case LTRD:
			crd = Coord{r + i, c + i}
		case RTLD:
			crd = Coord{r + i, c - i}
		case LTRU:
			crd = Coord{r - i, c + i}
		case RTLU:
			crd = Coord{r - i, c - i}
		}

		if crd.Row > int64(len(s)) || crd.Row < 0 || crd.Col > int64(len(s[crd.Row])) || crd.Col < 0 {
			return nil, nil, fmt.Errorf("Vector out of bounds [ROW|COL]:[%d|%d]", crd.Row, crd.Col)
		}

		if i < int64(len(n)) {
			s[crd.Row][crd.Col] = n[i]
		}

		res = append(res, s[crd.Row][crd.Col])
		crds = append(crds, crd)
	}

	return res, crds, nil
}
