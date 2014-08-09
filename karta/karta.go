package karta

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"code.google.com/p/draw2d/draw2d"
	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

var (
	green    = color.RGBA{0xD1, 0xE7, 0x51, 0xFF}
	black    = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white    = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	blue     = color.RGBA{0x4D, 0xBC, 0xE9, 0xFF}
	darkblue = color.RGBA{0x26, 0xAD, 0xE4, 0xFF}
	orange   = color.RGBA{0xFF, 0x66, 0x00, 0xff}
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

	draw.Draw(img, img.Bounds(), &image.Uniform{darkblue}, image.ZP, draw.Src)

	p := draw2d.NewGraphicContext(img)
	p.SetFillColor(green)

	l := draw2d.NewGraphicContext(img)
	l.SetLineWidth(3.0)
	l.SetStrokeColor(blue)

	// Iterate over cells
	for _, cell := range diagram.Cells {
		p.ArcTo(cell.Site.X, cell.Site.Y, 3, 3, 0, τ)
		p.FillStroke()

		for _, hedge := range cell.Halfedges {
			a := hedge.GetStartpoint()
			b := hedge.GetEndpoint()

			l.MoveTo(a.X, a.Y)
			l.LineTo(b.X, b.Y)
		}

		l.Stroke()
	}

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
