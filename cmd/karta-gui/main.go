package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
)

const qmlString = `
// Start of the QML string
import QtQuick 2.0
import QtQuick.Particles 2.0

Rectangle {
	id: root

	width: 800
	height: 600

	color: "#030f14"

	ParticleSystem {
		anchors.fill: parent

		ImageParticle {
			source: "image://star/FFFFFF88"

			rotation: 15
			rotationVariation: 45
			rotationVelocity: 35
			rotationVelocityVariation: 25
		}

		Emitter {
			anchors.fill: parent
			emitRate: 160
			lifeSpan: 2000
			lifeSpanVariation: 500

			size: 1
			endSize: 22
		}
	}

	Image {
		id: karta

		x: (parent.width - width)/2
		y: (parent.height - height)/2

		source: "image://karta/map.png"
	}
}

// End of the QML string`

func run() error {
	engine := qml.NewEngine()

	engine.AddImageProvider("karta", kartaImageProvider)
	engine.AddImageProvider("star", starImageProvider)

	component, err := engine.LoadString("karta.qml", qmlString)
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

// vim: ft=qml.go
