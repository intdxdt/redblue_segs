package redblue_segs

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
			if intersects(rseg[0], rseg[1], bseg[0], bseg[1]) {
				crossings = append(crossings, []int{i, j})
			}
		}
	}
	return crossings
}
