package redblue_segs

import (
	"sort"
	"github.com/intdxdt/math"
)

const (
	CreateRED  = iota
	CreateBLUE
	RemoveRED
	RemoveBLUE
)

type event struct {
	val float64
	ev  int
	idx int
}

//coordinates iterable of points
type events []*event

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
	var d = a.val - b.val
	var id int
	//x's are close enough to each other
	if math.FloatEqual(d, 0) {
		id = a.ev - b.ev
	} else {
		return d < 0
	}

	//y's are close enough to each other
	if id == 0 {
		id = a.idx - b.idx
	} else {
		return id < 0
	}

	return id < 0
}

//Lexicographic sort
func (o events) Sort() {
	sort.Sort(o)
}
