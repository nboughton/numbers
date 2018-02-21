/*
Package grid - functions for getting grid vectors, isiting vector paths and creating interesting grids
*/
package grid

import (
	"math"
)

/*
GetVector returns int64 slice of coord slices according to direction:

>: left to right

<: right to left

v: top to bottom

^: bottom to top

\v: left to right downwards diagonal

v/: right to left downwards diagonal

/^: left to right upwards diagonal

^\: right to left upwards diagonal

Where r, c are the starting point row and column, l is the
length of the vector and d is the symbol indicating direction
*/
func GetVector(r, c, l int64, d string) [][]int64 {
	v := make([][]int64, l)
	for i := int64(0); i < l; i++ {
		switch d {
		case `>`:
			v[i] = append(v[i], r, c+i)
		case `<`:
			v[i] = append(v[i], r, c-i)
		case `v`:
			v[i] = append(v[i], r+i, c)
		case `^`:
			v[i] = append(v[i], r-i, c)
		case `\v`:
			v[i] = append(v[i], r+i, c+i)
		case `v/`:
			v[i] = append(v[i], r+i, c-i)
		case `/^`:
			v[i] = append(v[i], r-i, c+i)
		case `^\`:
			v[i] = append(v[i], r-i, c-i)
		}
	}
	return v
}

// SafeVector returns false if a vector goes into negatives or over the
// limit (lim)
func SafeVector(v [][]int64, lim int64) bool {
	for _, cs := range v {
		for _, c := range cs {
			if c < 0 || c >= lim {
				return false
			}
		}
	}
	return true
}

/*
	Consider the following:
	[]int{
		[]int{25,10,11,12,13},
		[]int{24,09,02,03,14},
		[]int{23,08,01,04,15},
		[]int{22,07,06,05,16},
		[]int{21,20,19,18,17}
	}
*/

func numberSpiral(width int) [][]int {
	if width%2 == 0 {
		width++
	}

	grid := make([][]int, width)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	r, c := width/2, width/2
	grid[r][c] = 1
	for i := 1; i <= int(math.Pow(float64(width), 3)); i++ {

	}

	return grid
}

// CreateNumberSpiral creates a grid that contains a number spiral, n must
// be odd. If not it is incremented to become odd
func CreateNumberSpiral(n int64) [][]int64 {
	if n%2 == 0 {
		n++
	}

	// Create our array in memory
	g := make([][]int64, n)
	for i := range g {
		g[i] = make([]int64, n)
	}

	// Our starting point must be n/2, n/2
	r, c := int64(n/2), int64(n/2)
	g[r][c] = 1
	for i := int64(1); true; i++ {
		if i%2 != 0 {
			vct := GetVector(r, c, i+1, ">")
			// This is a get out for the last line of the spiral
			if !SafeVector(vct, n) {
				vct = GetVector(r, c, i, ">")
				numSpirIncrementVector(&g, vct)
				break
			}
			numSpirIncrementVector(&g, vct)
			numSpirSetRowCol(&r, &c, &vct)

			vct = GetVector(r, c, i+1, "v")
			numSpirIncrementVector(&g, vct)
			numSpirSetRowCol(&r, &c, &vct)
		} else {
			vct := GetVector(r, c, i+1, "<")
			numSpirIncrementVector(&g, vct)
			numSpirSetRowCol(&r, &c, &vct)

			vct = GetVector(r, c, i+1, "^")
			numSpirIncrementVector(&g, vct)
			numSpirSetRowCol(&r, &c, &vct)
		}
	}

	return g
}

func numSpirIncrementVector(g *[][]int64, vct [][]int64) {
	base := (*g)[vct[0][0]][vct[0][1]]
	for i, c := range vct {
		(*g)[c[0]][c[1]] = base + int64(i)
	}
}

func numSpirSetRowCol(r, c *int64, vct *[][]int64) {
	crd := (*vct)[len(*vct)-1]
	*r, *c = crd[0], crd[1]
}
