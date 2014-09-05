package main

const kartaQMLString = `
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

// vim: ft=qml.go
