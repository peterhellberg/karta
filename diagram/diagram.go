package diagram

import (
	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

// Diagram wraps a Voronoi diagram and contains
// a vertex pointing to the center of the map
type Diagram struct {
	*voronoi.Diagram
	Center voronoi.Vertex
}

// New generates a Voronoi diagram, relaxed by Lloyd's algorithm
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

	center := voronoi.Vertex{
		X: float64(w / 2),
		Y: float64(h / 2),
	}

	return &Diagram{d, center}
}

// Distance returns the distance between two vertices
func Distance(a, b voronoi.Vertex) float64 {
	return utils.Distance(a, b)
}
