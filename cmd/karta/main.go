package main

import (
	"flag"
	"log"
	"math/rand"
	"os/exec"

	"github.com/peterhellberg/karta"
)

var (
	output = flag.String("output", "karta.png", "Output filename")
	count  = flag.Int("count", 256, "The number of sites in the voronoi diagram")
	width  = flag.Int("width", 512, "The width of the map in pixels")
	height = flag.Int("height", 512, "The height of the map in pixels")
	relax  = flag.Int("iterations", 1, "The number of iterations of Lloyd's algorithm to run (max 16)")
	seed   = flag.Int64("seed", 7, "The starting seed for the map generator")
	show   = flag.Bool("show", false, "Show generated map using Preview.app")
)

func main() {
	flag.Parse()

	rand.Seed(*seed)

	diagram := karta.NewDiagram(float64(*width), float64(*height), *count, *relax)
	drawing := karta.DrawDiagramImage(diagram, *width, *height)

	err := karta.SaveImage(drawing, *output)
	if err != nil {
		log.Fatal(err)
	} else if *show {
		previewImage(*output)
	}
}

func previewImage(name string) {
	cmd := exec.Command("open", "-a", "/Applications/Preview.app", name)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
