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

// Cells represents a list of cells
type Cells []*Cell

// Karta represents the entire map
type Karta struct {
	Width   int
	Height  int
	Unit    float64
	Cells   Cells
	Diagram *diagram.Diagram
	Noise   *noise.Noise
	Image   image.Image
}

// New instantiates a new Karta
func New(w, h, c, r int) *Karta {
	return &Karta{
		Width:   w,
		Height:  h,
		Unit:    float64(math.Min(float64(w), float64(h)) / 20),
		Cells:   Cells{},
		Diagram: diagram.New(float64(w), float64(h), c, r),
		Noise:   noise.New(rand.Int63n(int64(w * h))),
	}
}

// Image creates a new Karta, then generates an Image
func Image(w, h, c, r int) image.Image {
	return New(w, h, c, r).GenerateImage()
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
		Cells  Cells   `json:"cells"`
	}{
		Width:  k.Width,
		Height: k.Height,
		Unit:   k.Unit,
		Edges:  edges,
		Cells:  k.Cells,
	}, "", "  ")
}

// GenerateImage generates an image
func (k *Karta) GenerateImage() image.Image {
	if k.Generate() == nil {
		return k.Image
	}

	return image.NewRGBA(image.Rect(0, 0, k.Width, k.Height))
}
