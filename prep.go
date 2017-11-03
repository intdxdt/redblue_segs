package redblue_segs

import "github.com/intdxdt/math"

func prepareEvts(red, blue [][][]float64) [][]float64 {
	var nr = len(red)
	var nb = len(blue)
	var n = nr + nb
	//var cols = 3
	//var list = ndscratch.malloc([2 * n, 3])
	var data = make([][]float64, 2*n)
	//var data = list.data
	var i, ptr int
	var x, y float64
	var seg [][]float64

	for i = 0; i < nr; i++ {
		seg = red[i]
		x, y = seg[0][0], seg[1][0]

		data[ptr] = []float64{math.MinF64(x, y), CREATE_RED, float64(i)}
		ptr += 1

		data[ptr] = []float64{math.MaxF64(x, y), REMOVE_RED, float64(i)}
		ptr += 1
	}
	for i = 0; i < nb; i++ {
		seg = blue[i]
		x, y = seg[0][0], seg[1][0]

		data[ptr] = []float64{math.MinF64(x, y), CREATE_BLUE, float64(i)}
		ptr += 1

		data[ptr] = []float64{math.MaxF64(x, y), REMOVE_BLUE, float64(i)}
		ptr += 1
	}
	events(data).Sort()
	return data
}
