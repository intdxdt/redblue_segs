package redblue_segs

import (
	"fmt"
	"time"
	"sort"
	"testing"
	"math/rand"
	rbt "github.com/intdxdt/redblue_segs/test"
	"github.com/franela/goblin"
)

var random = rand.Float64

func TestFuzz(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("fuzz test", func() {
		g.It("should test fuzzy", func() {
			g.Timeout(1 * time.Hour)

			for j := 0; j < 10; j++ {
				fmt.Println("# fuzz test ", j+1, " ...")
				var red = make([][][]float64, 0)
				for i := 0; i < 10*(j+1); i++ {
					red = append(red, [][]float64{
						{random(), random()}, {random(), random()},
					})
				}

				var blue = make([][][]float64, 0)
				for i := 0; i < 10*(j+1); i++ {
					blue = append(blue, [][]float64{
						{random(), random()},
						{random(), random()},
					})
				}

				var expected = brutal(red, blue)
				sort.Sort(rbt.Crossings(expected))

				var actual = rblsi(red, blue)
				sort.Sort(rbt.Crossings(actual))
				g.Assert(actual).Equal(expected)
			}
		})
	})
}

