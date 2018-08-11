package redblue_segs

import (
	"github.com/intdxdt/math"
)

func addSegment(
	index int, red [][][]float64, redList *bruteForceList,
	blue [][][]float64, blueList *bruteForceList,
	visit func(int, int) bool, flip bool) bool {
	//Look up segment
	var seg = red[index]
	var bseg [][]float64
    //Get segment end points
    var segA = seg[0]
    var segB = seg[1]

	//Read out components
	var a0 = segA[1]
	var a1 = segB[1]
	var l0 = math.MinF64(a0, a1)
	var h0 = math.MaxF64(a0, a1)

	//Scan over blue intervals for point
	var count = blueList.count
	var ptr = 2 * count
	var h1, l1 float64
	var bindex int
	var ret bool


	for i := count - 1;!ret && i >= 0; i-- {
		ptr += -1
		h1 = blueList.intervals[ptr]
		ptr += -1
		l1 = blueList.intervals[ptr]

		//Test if intervals overlap
		if l0 <= h1 && l1 <= h0 {
			bindex = blueList.index[i]
			bseg = blue[bindex]

			//Test if segments intersect
			if intersects(segA, segB, bseg[0], bseg[1]) {
				if flip {
					ret = visit(bindex, index)
				} else {
					ret = visit(index, bindex)
				}
			}
		}
	}

	redList.insert(l0, h0, index)
	return false
}
