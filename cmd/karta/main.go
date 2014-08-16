package main

import (
	"flag"
	"image/png"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"

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

	if *count < 3 {
		log.Fatalf("count must be at least 3")
	}

	// Create a new karta
	k := karta.New(*width, *height, *count, *relax)

	if k.Generate() == nil {
		file, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if strings.HasSuffix(*output, ".json") {
			if j, err := k.MarshalJSON(); err == nil {
				file.Write(j)
			}
		} else {
			err := png.Encode(file, k.Image)

			if err != nil {
				log.Fatal(err)
			} else if *show {
				previewImage(*output)
			}
		}
	}
}

func previewImage(name string) {
	cmd := exec.Command("open", "-a", "/Applications/Preview.app", name)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
