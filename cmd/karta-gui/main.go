package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

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

	ctrl := Control{Message: "Hello from Go!"}

	context := engine.Context()
	context.SetVar("ctrl", &ctrl)

	win := component.CreateWindow(nil)

	ctrl.Root = win.Root()

	win.Show()
	win.Wait()

	return nil
}

func main() {
	// Seed the random number generator
	rand.Seed(int64(time.Now().Second()))

	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type Control struct {
	Root    qml.Object
	Message string
}

func (ctrl *Control) TextReleased(text qml.Object) {
	ctrl.Message = text.String("text")

	qml.Changed(ctrl, &ctrl.Message)
}
