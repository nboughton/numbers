package set

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

// Int64s is a slice of slices of int64
type Int64s []Int64

func (s Int64s) Len() int           { return len(s) }
func (s Int64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int64s) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

// Find returns all sets in s that contain n
func (s Int64s) Find(n int64) Int64s {
	var res Int64s

	for _, set := range s {
		if set.Contains(n) {
			res = append(res, set)
		}
	}

	return res
}
