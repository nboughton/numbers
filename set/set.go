package set

// Int64 is a slice of int64
type Int64 []int64

func (s Int64) Len() int           { return len(s) }
func (s Int64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int64) Less(i, j int) bool { return s[i] < s[j] }

// Int64s is a slice of slices of int64
type Int64s []Int64

func (s Int64s) Len() int           { return len(s) }
func (s Int64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int64s) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
