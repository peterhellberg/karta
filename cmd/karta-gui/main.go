package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
)

func run() error {
	engine := qml.NewEngine()

	engine.AddImageProvider("karta", kartaImageProvider)
	engine.AddImageProvider("star", starImageProvider)

	component, err := engine.LoadString("karta.qml", kartaQMLString)
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)

	win.Show()
	win.Wait()

	return nil
}

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
