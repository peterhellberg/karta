package karta

import "image/color"

type Cell struct {
	Index          int
	CenterDistance float64
	Noise          float64
	Elevation      float64
	Land           bool
	FillColor      color.RGBA
	StrokeColor    color.RGBA
}
