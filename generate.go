package karta

import (
	"image"
	"math/rand"

	"github.com/peterhellberg/karta/diagram"
	. "github.com/peterhellberg/karta/palette"

	"code.google.com/p/draw2d/draw2d"
)

func (k *Karta) Generate() error {
	k.generateGeography()
	k.drawImage()

	return nil
}

func (k *Karta) generateGeography() {
	w := k.Width
	h := k.Height
	u := k.Unit

	for i, cell := range k.Diagram.Cells {
		d := diagram.Distance(cell.Site, k.Diagram.Center)
		n := k.Noise.Noise2D(
			cell.Site.X/(float64(k.Width)/3.7),
			cell.Site.Y/(float64(k.Height)/3.7))

		e := elevation(k, d, n)
		c := k.SetCell(i, &Cell{
			Index:          i,
			CenterDistance: d,
			Noise:          n,
			Elevation:      e,
			Land:           e >= 0,
		})

		// Make sure left and right edges of the map are deep water
		if cell.Site.X < u*0.5 || cell.Site.X > float64(w)-u*0.5 {
			c.Elevation -= 1.2
		}

		// Make sure top and bottom edges of the map are deep water
		if cell.Site.Y < u/1.5 || cell.Site.Y > float64(h)-u/1.5 {
			c.Elevation -= 1.2
		}

		if cell.Site.Y < u/3 || cell.Site.Y > float64(h)-u/3 {
			c.Elevation -= 1.1
		}

		if c.Land {
			c.Elevation -= c.Noise * 1.6

			if d < u*3.3 {
				c.Elevation += 0.3
			}

			if d < u*2.3 {
				c.Elevation += 0.6
			}

			if d < u*1.3 {
				c.Elevation += 0.9
			}

			switch {
			case c.Elevation > 7:
				c.FillColor = Green7
				c.StrokeColor = Green8
			case c.Elevation > 6.1:
				c.FillColor = Green6
				c.StrokeColor = Green7
			case c.Elevation > 4.8:
				c.FillColor = Green5
				c.StrokeColor = Green6
			case c.Elevation > 3.1:
				c.FillColor = Green4
				c.StrokeColor = Green5
			case c.Elevation > 2.4:
				c.FillColor = Green3
				c.StrokeColor = Green4
			case c.Elevation > 1.5:
				c.FillColor = Green2
				c.StrokeColor = Green3
			case c.Elevation < 0:
				c.FillColor = Yellow1
				c.StrokeColor = Yellow2
			default:
				c.FillColor = Green1
				c.StrokeColor = Green2
			}

			// Add some lakes
			switch {
			case c.Noise < -0.6:
				c.FillColor = Blue1
				c.StrokeColor = Blue2
			case c.Noise < -0.4:
				c.FillColor = Blue0
				c.StrokeColor = Blue1
			}
		} else {
			switch {
			case c.Elevation < -6:
				c.FillColor = Blue7
				c.StrokeColor = Blue7
			case c.Elevation < -5:
				c.FillColor = Blue6
				c.StrokeColor = Blue7
			case c.Elevation < -4:
				c.FillColor = Blue5
				c.StrokeColor = Blue6
			case c.Elevation < -3:
				c.FillColor = Blue4
				c.StrokeColor = Blue5
			case c.Elevation < -2:
				c.FillColor = Blue3
				c.StrokeColor = Blue4
			case c.Elevation < -1:
				c.FillColor = Blue2
				c.StrokeColor = Blue3
			default:
				c.FillColor = Blue1
				c.StrokeColor = Blue2
			}
		}
	}
}

func (k *Karta) drawImage() {
	img := image.NewRGBA(image.Rect(0, 0, k.Width, k.Height))

	l := draw2d.NewGraphicContext(img)

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
