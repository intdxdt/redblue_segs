package redblue_segs

import (
	"github.com/intdxdt/geom/segs"
)

func BruteForce(red, blue [][][]float64) [][]int {
	var nr = len(red)
	var nb = len(blue)
	var crossings = make([][]int, 0)

	for i := 0; i < nr; i++ {
		var rseg = red[i]
		//var a    = rseg[0]
		//var b    = rseg[1]
		for j := 0; j < nb; j++ {
			bseg := blue[j]
			//c := bseg[0]
			//d := bseg[1]
			if segs.Intersects(rseg, bseg) {
				crossings = append(crossings, []int{i, j})
			}
		}
	}
	return crossings
}
