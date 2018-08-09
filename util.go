package redblue_segs

func maxf64(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}

func minf64(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}
