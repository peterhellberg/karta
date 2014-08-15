package diagram

import (
	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

type Diagram struct {
	*voronoi.Diagram
	Center voronoi.Vertex
}

// Diagram generates a new Voronoi diagram, relaxed by Lloyd's algorithm
func New(w, h float64, c, r int) *Diagram {
	bbox := voronoi.NewBBox(0, w, 0, h)
	sites := utils.RandomSites(bbox, c)

	// Compute voronoi diagram.
	d := voronoi.ComputeDiagram(sites, bbox, true)

	// Max number of iterations is 16
	if r > 16 {
		r = 16
	}

	// Relax using Lloyd's algorithm
	for i := 0; i < r; i++ {
		sites = utils.LloydRelaxation(d.Cells)
		d = voronoi.ComputeDiagram(sites, bbox, true)
	}

	center := voronoi.Vertex{float64(w / 2), float64(h / 2)}

	return &Diagram{d, center}
}

func Distance(a, b voronoi.Vertex) float64 {
	return utils.Distance(a, b)
}
