package karta

import (
	"encoding/json"
	"image"
	"math"
	"math/rand"

	"github.com/peterhellberg/karta/diagram"
	"github.com/peterhellberg/karta/noise"
	"github.com/pzsz/voronoi"
)

type Karta struct {
	Width   int
	Height  int
	Unit    float64
	Cells   []*Cell
	Diagram *diagram.Diagram
	Noise   *noise.Noise
	Image   image.Image
}

func New(w, h, c, r int) *Karta {
	return &Karta{
		Width:   w,
		Height:  h,
		Unit:    float64(math.Min(float64(w), float64(h)) / 20),
		Cells:   []*Cell{},
		Diagram: diagram.New(float64(w), float64(h), c, r),
		Noise:   noise.New(rand.Int63n(int64(w * h))),
	}
}

// Edge represents two vertexes
type Edge struct {
	VaVertex voronoi.Vertex
	VbVertex voronoi.Vertex
}

// MarshalJSON marshals the map as JSON
func (k *Karta) MarshalJSON() ([]byte, error) {
	edges := []*Edge{}

	for _, ev := range k.Diagram.Edges {
		edges = append(edges, &Edge{
			VaVertex: ev.Va.Vertex,
			VbVertex: ev.Vb.Vertex,
		})
	}

	return json.MarshalIndent(struct {
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Unit   float64 `json:"unit"`
		Edges  []*Edge `json:"edges"`
		Cells  []*Cell `json:"cells"`
	}{
		Width:  k.Width,
		Height: k.Height,
		Unit:   k.Unit,
		Edges:  edges,
		Cells:  k.Cells,
	}, "", "  ")
}
