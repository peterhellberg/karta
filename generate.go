package karta

import (
	"image"
	"math/rand"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/peterhellberg/karta/diagram"
	"github.com/peterhellberg/karta/palette"
)

// Generate generates a map
func (k *Karta) Generate() error {
	k.generateTopography()
	k.drawImage()

	return nil
}

func (k *Karta) generateTopography() {
	u := k.Unit

	for i, cell := range k.Diagram.Cells {
		d := diagram.Distance(cell.Site, k.Diagram.Center)
		n := k.Noise.Noise2D(
			cell.Site.X/(float64(k.Width)/4),
			cell.Site.Y/(float64(k.Height)/4))

		e := elevation(k, d, n)
		c := &Cell{
			Index:          i,
			CenterDistance: d,
			NoiseLevel:     n,
			Elevation:      e,
			Land:           e >= 0,
			Site:           cell.Site,
		}

		if c.Land {
			// Make sure edges of the map are water
			if (cell.Site.X < u*0.5 || cell.Site.X > float64(k.Width)-u*0.5) ||
				(cell.Site.Y < u/1.5 || cell.Site.Y > float64(k.Height)-u/1.5) ||
				(cell.Site.Y < u/3 || cell.Site.Y > float64(k.Height)-u/3) {
				c.Land = false
				c.Elevation = -1.5 * c.NoiseLevel
			}
		}

		if c.Land {
			if d < u*3.3 {
				c.Elevation += 0.3
			}

			if d < u*2.3 {
				c.Elevation += 0.6
			}

			if d < u*1.3 {
				c.Elevation += 0.9
			}

			// Add some lakes
			if c.NoiseLevel < -0.3 {
				c.Elevation = c.NoiseLevel
			}

			switch {
			case c.Elevation > 7:
				c.FillColor = palette.Green7
				c.StrokeColor = palette.Green8
			case c.Elevation > 6.1:
				c.FillColor = palette.Green6
				c.StrokeColor = palette.Green7
			case c.Elevation > 4.8:
				c.FillColor = palette.Green5
				c.StrokeColor = palette.Green6
			case c.Elevation > 3.1:
				c.FillColor = palette.Green4
				c.StrokeColor = palette.Green5
			case c.Elevation > 2.4:
				c.FillColor = palette.Green3
				c.StrokeColor = palette.Green4
			case c.Elevation > 1.5:
				c.FillColor = palette.Green2
				c.StrokeColor = palette.Green3
			case c.Elevation < -0.6:
				c.FillColor = palette.Blue1
				c.StrokeColor = palette.Blue2
			case c.Elevation < -0.4:
				c.FillColor = palette.Blue0
				c.StrokeColor = palette.Blue1
			case c.Elevation < 0:
				c.FillColor = palette.Yellow1
				c.StrokeColor = palette.Yellow2
			default:
				c.FillColor = palette.Green1
				c.StrokeColor = palette.Green2
			}
		} else {
			switch {
			case c.Elevation < -6:
				c.FillColor = palette.Blue7
				c.StrokeColor = palette.Blue7
			case c.Elevation < -5:
				c.FillColor = palette.Blue6
				c.StrokeColor = palette.Blue7
			case c.Elevation < -4:
				c.FillColor = palette.Blue5
				c.StrokeColor = palette.Blue6
			case c.Elevation < -3:
				c.FillColor = palette.Blue4
				c.StrokeColor = palette.Blue5
			case c.Elevation < -2:
				c.FillColor = palette.Blue3
				c.StrokeColor = palette.Blue4
			case c.Elevation < -1:
				c.FillColor = palette.Blue2
				c.StrokeColor = palette.Blue3
			default:
				c.FillColor = palette.Blue1
				c.StrokeColor = palette.Blue2
			}
		}

		k.Cells = append(k.Cells, c)
	}
}

func (k *Karta) drawImage() {
	img := image.NewRGBA(image.Rect(0, 0, k.Width, k.Height))

	l := draw2dimg.NewGraphicContext(img)

	l.SetLineWidth(1.2)

	// Iterate over cells
	for i, cell := range k.Diagram.Cells {
		l.SetFillColor(k.Cells[i].FillColor)
		l.SetStrokeColor(k.Cells[i].StrokeColor)

		for _, hedge := range cell.Halfedges {
			a := hedge.GetStartpoint()
			b := hedge.GetEndpoint()

			l.MoveTo(a.X, a.Y)
			l.LineTo(b.X, b.Y)
		}

		l.FillStroke()
	}

	l.Close()

	k.Image = img
}

func elevation(k *Karta, d, n float64) (e float64) {
	e = 1.8 + n

	e -= (d / k.Unit) / 3.75

	if e > 0 {
		e += 1 + float64(rand.Int63n(2))

		if e > 1.5 && rand.Intn(3) < 2 {
			e += 0.5 + rand.Float64()
		}

		if e > 3 {
			e += 1.5 + rand.Float64()
		}
	}

	return
}
