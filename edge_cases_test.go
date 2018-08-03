package redblue_segs

import (
	"time"
	"testing"
	"github.com/franela/goblin"
)

func TestEdgeCase(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("vertical edge case", func() {
		g.It("should test edge case", func() {
			g.Timeout(1 * time.Hour)
              var red  = [][][]float64{{ { 224, 328 }, { 224, 331 }} }
              var blue = [][][]float64{{ { 224, 146 }, { 224, 330 }} }
              var visit = func  (r, b int ) bool{
                  g.Assert(true).IsTrue()
                  return false
              }
              RedBlueLineSegmentIntersection(red, blue, visit)
		})
	})
}

