package karta

import (
	"image"
	"math"
	"math/rand"

	"github.com/peterhellberg/karta/diagram"
	"github.com/peterhellberg/karta/noise"
)

type Karta struct {
	Width   int
	Height  int
	Unit    float64
	Cells   map[int]*Cell
	Diagram *diagram.Diagram
	Noise   *noise.Noise
	Image   image.Image
}

func New(w, h, c, r int) *Karta {
	return &Karta{
		Width:   w,
		Height:  h,
		Unit:    float64(math.Min(float64(w), float64(h)) / 20),
		Cells:   make(map[int]*Cell),
		Diagram: diagram.New(float64(w), float64(h), c, r),
		Noise:   noise.New(rand.Int63n(int64(w * h))),
	}
}

func (k *Karta) SetCell(i int, c *Cell) *Cell {
	k.Cells[i] = c

	return c
}
