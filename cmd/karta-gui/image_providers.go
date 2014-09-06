package main

import (
	"image"
	"image/color"

	"github.com/peterhellberg/karta"
)

func kartaImageProvider(name string, width, height int) image.Image {
	if width == 0 {
		width = 512
	}

	if height == 0 {
		height = 512
	}

	count := 512
	iterations := 1

	// Create a new karta
	k := karta.New(width, height, count, iterations)

	if k.Generate() == nil {
		return k.Image
	}

	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func starImageProvider(name string, width, height int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 11, 11))

	c := hexToColorNRGBA(name)

	r := c.R
	g := c.G
	b := c.B
	a := c.A

	au := a / 5

	for x := 0; x <= 11; x++ {
		if x < 6 {
			img.Set(x, 5, color.NRGBA{r, g, b, au * uint8(x)})
		}

		if x > 5 {
			img.Set(x, 5, color.NRGBA{r, g, b, au * (uint8(5) - uint8(-6+x))})
		}
	}

	img.Set(5, 5, c)

	for y := 0; y <= 11; y++ {
		if y < 6 {
			img.Set(5, y, color.NRGBA{r, g, b, au * uint8(y)})
		}

		if y > 5 {
			img.Set(5, y, color.NRGBA{r, g, b, au * (uint8(5) - uint8(-6+y))})
		}
	}

	return img
}
