package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"

	"code.google.com/p/draw2d/draw2d"

	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

var (
	output = flag.String("output", "karta.png", "Output filename")
	count  = flag.Int("count", 256, "The number of sites in the voronoi diagram")
	width  = flag.Int("width", 512, "The width of the map in pixels")
	height = flag.Int("height", 512, "The height of the map in pixels")
	relax  = flag.Int("iterations", 1, "The number of iterations of Lloyd's algorithm to run (max 16)")
	seed   = flag.Int64("seed", 7, "The starting seed for the map generator")
	show   = flag.Bool("show", false, "Show generated map using Preview.app")

	green    = color.RGBA{0xD1, 0xE7, 0x51, 0xFF}
	black    = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white    = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	blue     = color.RGBA{0x4D, 0xBC, 0xE9, 0xFF}
	darkblue = color.RGBA{0x26, 0xAD, 0xE4, 0xFF}
	orange   = color.RGBA{0xFF, 0x66, 0x00, 0xff}
)

func main() {
	flag.Parse()

	rand.Seed(*seed)

	diagram := NewDiagram(float64(*width), float64(*height), *count, *relax)
	drawing := DrawDiagramImage(diagram, *width, *height)

	err := SaveImage(drawing, *output)
	if err != nil {
		log.Fatal(err)
	}
}

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

	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	p := draw2d.NewGraphicContext(img)
	p.SetFillColor(green)

	l := draw2d.NewGraphicContext(img)
	l.SetLineWidth(2.1)
	l.SetStrokeColor(blue)

	// Iterate over cells
	for _, cell := range diagram.Cells {
		p.ArcTo(cell.Site.X, cell.Site.Y, 2, 2, 0, 2*math.Pi)
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

	if *show {
		previewImage(file.Name())
	}

	return nil
}

func previewImage(name string) {
	cmd := exec.Command("open", "-a", "/Applications/Preview.app", name)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
