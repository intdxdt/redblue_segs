package redblue_segs

func prepareEvents(red, blue [][][]float64) []event {
	var nr = len(red)
	var nb = len(blue)
	var n = nr + nb
	var x, y float64
	var seg [][]float64
	var data = make([]event, 0, 2*n)

	for i := 0; i < nr; i++ {
		seg = red[i]
		x, y = seg[0][0], seg[1][0]
		data  =append(data,  event{minf64(x, y), CreateRED, i})
		data = append(data,  event{maxf64(x, y), RemoveRED, i})
	}

	for i := 0; i < nb; i++ {
		seg = blue[i]
		x, y = seg[0][0], seg[1][0]
		data = append(data,  event{minf64(x, y), CreateBLUE, i})
		data = append(data,  event{maxf64(x, y), RemoveBLUE, i})
	}
	events(data).Sort()
	return data
}
