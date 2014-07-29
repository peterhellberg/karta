package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

var (
	count  = flag.Int("c", 256, "The number of sites in the voronoi diagram")
	width  = flag.Int("w", 512, "The width of the map in pixels")
	height = flag.Int("h", 512, "The height of the map in pixels")
	seed   = flag.Int64("s", 7, "The starting seed for the map generator")
	output = flag.String("o", "karta.png", "Output filename")
	show   = flag.Bool("show", true, "Show generated map using Preview.app")

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

	bbox := voronoi.NewBBox(0, float64(*width), 0, float64(*height))

	sites := utils.RandomSites(bbox, *count)

	// Compute voronoi diagram.
	diagram := voronoi.ComputeDiagram(sites, bbox, true)

	img := image.NewRGBA(image.Rect(0, 0, *width, *height))

	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	// Iterate over cells
	for _, cell := range diagram.Cells {
		x := int(cell.Site.X)
		y := int(cell.Site.Y)

		img.Set(x, y, orange)
	}

	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	png.Encode(file, img)

	if *show {
		Show(file.Name())
	}
}

func Show(name string) {
	cmd := exec.Command("open", "-a", "/Applications/Preview.app", name)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
