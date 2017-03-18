package main

import (
	"flag"
	"math"
	"math/rand"

	"github.com/peterhellberg/karta"

	. "github.com/fogleman/fauxgl"
)

const (
	fovy  = 10
	near  = 1
	far   = 100
	relax = 4
	count = 2048
)

var (
	eye    = V(8, 5, 5)
	center = V(0.04, 0, 0)
	up     = V(0, 0, 1)

	coal = HexColor("191616")
)

func main() {
	var seed int64

	flag.Int64Var(&seed, "seed", 4, "The starting seed for the map generator")

	flag.Parse()

	rand.Seed(seed)

	k := karta.New(256, 256, count, relax)

	if k.Generate() == nil {
		voxels := []Voxel{}

		m := k.Image

		w, h := m.Bounds().Max.X, m.Bounds().Max.Y

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				z := getZ(k.Cells, x, y)

				for i := 0; i < z; i++ {
					voxels = append(voxels, Voxel{x, y, i + 1, MakeColor(m.At(x, y))})
				}
			}
		}

		mesh := NewVoxelMesh(voxels)

		mesh.BiUnitCube()

		context := NewContext(1024, 768)
		context.ClearColor = coal
		context.ClearColorBuffer()

		aspect := float64(1024) / float64(768)
		matrix := LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
		light := V(1, 0.6, 1.9).Normalize()

		shader := NewPhongShader(matrix, light, eye)

		shader.DiffuseColor = Gray(1.2)

		context.Shader = shader
		context.DrawMesh(mesh)

		SavePNG("karta.png", context.Image())
	}
}

func getZ(cells karta.Cells, x, y int) int {
	sd := math.Inf(1)

	var cc *karta.Cell

	for _, c := range cells {
		if d := distance(float64(x), float64(y), c.Site.X, c.Site.Y); d < sd {
			sd = d
			cc = c
		}
	}

	if cc != nil {
		return 2 + int(cc.Elevation)
	}

	return 2
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
