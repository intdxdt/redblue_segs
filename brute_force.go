package redblue_segs

func BruteForce(red, blue [][][]float64) [][]int {
	var nr = len(red)
	var nb = len(blue)
	var crossings [][]int
	var rseg , bseg [][]float64

	for i := 0; i < nr; i++ {
		rseg = red[i]
		for j := 0; j < nb; j++ {
			bseg = blue[j]
			if intersects(rseg[0], rseg[1], bseg[0], bseg[1]) {
				crossings = append(crossings, []int{i, j})
			}
		}
	}
	return crossings
}
