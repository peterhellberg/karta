package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"github.com/peterhellberg/karta"
)

var (
	output = flag.String("output", "karta.png", "Output filename")
	count  = flag.Int("count", 2048, "The number of sites in the voronoi diagram")
	width  = flag.Int("width", 512, "The width of the map in pixels")
	height = flag.Int("height", 512, "The height of the map in pixels")
	relax  = flag.Int("iterations", 1, "The number of iterations of Lloyd's algorithm to run (max 16)")
	seed   = flag.Int64("seed", 3, "The starting seed for the map generator")
	show   = flag.Bool("show", false, "Show generated map using Preview.app")
)

func main() {
	// Parse the command line flags
	flag.Parse()

	// Seed the random number generator
	rand.Seed(*seed)

	// Create a new karta
	k := karta.New(*width, *height, *count, *relax)

	if k.Generate() == nil {
		err := saveImage(k.Image, *output)
		if err != nil {
			log.Fatal(err)
		} else if *show {
			previewImage(*output)
		}
	}
}

// saveImage saves an image to a file
func saveImage(img image.Image, fn string) error {
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

func previewImage(name string) {
	cmd := exec.Command("open", "-a", "/Applications/Preview.app", name)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
