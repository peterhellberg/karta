package main

import (
	"fmt"
	"image"
	"math/rand"
	"os"

	"github.com/peterhellberg/karta"
	"gopkg.in/qml.v1"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func kartaImageProvider(id string, width, height int) image.Image {
	if width == 0 {
		width = 512
	}

	if height == 0 {
		height = 512
	}

	seed := 3
	count := 512
	iterations := 1

	// Seed the random number generator
	rand.Seed(int64(seed))

	// Create a new karta
	k := karta.New(width, height, count, iterations)

	if k.Generate() == nil {
		return k.Image
	}

	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func run() error {
	engine := qml.NewEngine()
	engine.AddImageProvider("karta", kartaImageProvider)

	component, err := engine.LoadString("karta-gui", qmlString)
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)

	win.Show()
	win.Wait()

	return nil
}

const qmlString = `
import QtQuick 2.0

Image {
	source: "image://karta/map.png"
}
`
