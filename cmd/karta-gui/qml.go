package main

const kartaQMLString = `
// Start of the QML string
import QtQuick 2.0
import QtQuick.Particles 2.0
import GoExtensions 1.0

Rectangle {
	id: root

	width: 800
	height: 600

	color: "#030f14"

	property alias seed: seedInput.text

	Rectangle {
    id: form
    width: parent.width; height: 50

		y: parent.height-50

		color: "#061f29"

    Row {
      id: row
      anchors.centerIn: parent
			spacing: 20

			Text {
				id: helloText
				text: "seed:"

				color: "#FFFFFF"
				font.pointSize: 12
				font.bold: true
			}

			TextInput {
				id: seedInput

				color: "#1e8bb8"
				font.pointSize: 12
				font.bold: true

				width: 96;
				height: 20
				focus: true
				text: "1"
			}
    }
	}

	Rectangle {
		id: background

		width: parent.width; height: parent.height-50

		gradient: Gradient {
  		GradientStop { position: 0.0; color: "#061f29" }
    	GradientStop { position: 1.0; color: "#030f14" }
		}

		MouseArea {
  		id: area
  	  width: parent.width
  	  height: parent.height
			onClicked: {
				ip.visible = !ip.visible
			}
		}

		ParticleSystem {
			anchors.fill: parent

			ImageParticle {
				id: ip

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

			source: "image://karta/map.png"

			x: (parent.width - width)/2
			y: (parent.height - height)/2

			width: parent.width/2
			height: parent.height/2

			fillMode: Image.PreserveAspectCrop
			clip: true
		}

		GoRect {
			x: 60; y: 60; width: 200; height: 150
			SequentialAnimation on x {
				loops: Animation.Infinite

				NumberAnimation { from: 60; to: 320; duration: 4000; easing.type: Easing.InOutQuad }
				NumberAnimation { from: 320; to: 60; duration: 4000; easing.type: Easing.InOutQuad }
			}
		}
	}
}

// End of the QML string`

// vim: ft=qml.go
