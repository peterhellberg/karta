package karta

import (
	"image/color"

	"github.com/pzsz/voronoi"
)

// Cell is the smalles unit on the map
type Cell struct {
	Index          int            `json:"index"`
	CenterDistance float64        `json:"center_distance"`
	NoiseLevel     float64        `json:"noise_level"`
	Elevation      float64        `json:"elevation"`
	Land           bool           `json:"land"`
	Site           voronoi.Vertex `json:"site"`
	FillColor      color.RGBA     `json:"fill_color"`
	StrokeColor    color.RGBA     `json:"stroke_color"`
}
