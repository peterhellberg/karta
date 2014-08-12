package karta

import (
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"

	"github.com/peterhellberg/karta/palette"

	"code.google.com/p/draw2d/draw2d"
	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

const (
	// Define Tau since the Golang math package is lacking
	τ = 2 * math.Pi
)

// NewDiagram generates a new Voronoi diagram, relaxed by Lloyd's algorithm
func NewDiagram(w, h float64, c, r int) *voronoi.Diagram {
	bbox := voronoi.NewBBox(0, w, 0, h)
	sites := utils.RandomSites(bbox, c)

	// Compute voronoi diagram.
	diagram := voronoi.ComputeDiagram(sites, bbox, true)

	// Max number of iterations is 16
	if r > 16 {
		r = 16
	}

	// Relax using Lloyd's algorithm
	for i := 0; i < r; i++ {
		sites = utils.LloydRelaxation(diagram.Cells)
		diagram = voronoi.ComputeDiagram(sites, bbox, true)
	}

	return diagram
}

// DrawDiagramImage draws a Voroni diagram to an image
func DrawDiagramImage(diagram *voronoi.Diagram, w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	draw.Draw(img, img.Bounds(), &image.Uniform{palette.Darkblue}, image.ZP, draw.Src)

	p := draw2d.NewGraphicContext(img)
	p.SetFillColor(palette.Green)

	l := draw2d.NewGraphicContext(img)
	l.SetLineWidth(2)
	l.SetStrokeColor(palette.Purple)

	// Iterate over cells
	for _, cell := range diagram.Cells {
		l.SetFillColor(palette.Pink)

		for _, hedge := range cell.Halfedges {
			a := hedge.GetStartpoint()
			b := hedge.GetEndpoint()

			l.MoveTo(a.X, a.Y)
			l.LineTo(b.X, b.Y)
		}

		l.FillStroke()

		p.ArcTo(cell.Site.X, cell.Site.Y, 3, 3, 0, τ)
		p.FillStroke()
	}

	l.Close()

	return img
}

// SaveImage saves an image to a file
func SaveImage(img image.Image, fn string) error {
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
