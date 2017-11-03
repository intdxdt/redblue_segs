package redblue_segs

import (
	"sort"
	"github.com/intdxdt/math"
)

const (
	CREATE_RED  = iota
	CREATE_BLUE
	REMOVE_RED
	REMOVE_BLUE
)


const (
	x = iota
	y
	z
)

//coordinates iterable of points
type events [][]float64

//Len for sort interface
func (o events) Len() int {
	return len(o)
}

//Swap for sort interface
func (o events) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

//lexical sort of x and y coordinates
func (o events) Less(i, j int) bool {
	var a, b = o[i], o[j]
	var d = a[x] - b[x]

	//x's are close enough to each other
	if math.FloatEqual(d, 0) {
		d = a[y] - b[y]
	} else {
		return d < 0
	}

	//y's are close enough to each other
	if math.FloatEqual(d, 0) {
		d = a[z] - b[z]
	} else {
		return d < 0
	}

	return d < 0
}

//Inplace Lexicographic sort
func (o events) Sort() {
	sort.Sort(o)
}
