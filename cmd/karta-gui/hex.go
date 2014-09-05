package main

import (
	"image/color"
	"strconv"
)

func hexToUint8(hex string) (uint8, error) {
	i, err := strconv.ParseUint(hex, 16, 0)

	return uint8(i), err
}

func hexToColorNRGBA(hex string) color.NRGBA {
	c := color.NRGBA{}

	getIntensity := func(str string) uint8 {
		if u, err := hexToUint8(str); err == nil {
			return u
		}

		return 0
	}

	if len(hex) >= 6 {
		c.R = getIntensity(hex[0:2])
		c.G = getIntensity(hex[2:4])
		c.B = getIntensity(hex[4:6])
	}

	if len(hex) == 8 {
		c.A = getIntensity(hex[6:8])
	} else {
		c.A = 255
	}

	return c
}
