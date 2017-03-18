package main

import (
	"flag"
	"log"
	"math/rand"

	"github.com/peterhellberg/karta"

	. "github.com/fogleman/fauxgl"
)

var (
	output = flag.String("output", "karta.png", "Output filename")
	count  = flag.Int("count", 1024, "The number of sites in the voronoi diagram")
	width  = flag.Int("width", 128, "The width of the map in pixels")
	height = flag.Int("height", 128, "The height of the map in pixels")
	relax  = flag.Int("iterations", 4, "The number of iterations of Lloyd's algorithm to run (max 16)")
	seed   = flag.Int64("seed", 4, "The starting seed for the map generator")
)

const (
	fovy = 16
	near = 1
	far  = 40
)

var (
	eye    = V(8, 5, 5)
	center = V(0.04, 0, 0)
	up     = V(0, 0, 1)

	lime   = HexColor("D1F2A5")
	pink   = HexColor("F56991")
	purple = HexColor("551033")
	coal   = HexColor("191616")
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
		m := k.Image

		w, h := m.Bounds().Max.X, m.Bounds().Max.Y

		voxels := []Voxel{}

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				pc := m.At(x, y)

				_, _, b, _ := pc.RGBA()

				levels := 3

				if b > 30000 {
					levels = 2
				} else if b > 20000 {
					levels = 1
				}

				for z := 0; z < levels; z++ {
					voxels = append(voxels, Voxel{x, y, z, MakeColor(pc)})
				}
			}
		}

		mesh := NewVoxelMesh(voxels)

		mesh.BiUnitCube()

		context := NewContext(2048, 2048)
		context.ClearColor = coal
		context.ClearColorBuffer()

		aspect := float64(2048) / float64(2048)
		matrix := LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
		light := V(1, 0.6, 1.9).Normalize()

		shader := NewPhongShader(matrix, light, eye)

		shader.DiffuseColor = Gray(1.2)

		context.Shader = shader
		context.DrawMesh(mesh)

		SavePNG(*output, context.Image())
	}
}
