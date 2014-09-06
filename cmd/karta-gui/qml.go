package main

const kartaQMLString = `
// Start of the QML string
import QtQuick 2.0
import QtQuick.Particles 2.0
import GoExtensions 1.0

Rectangle {
	id: root

	property alias seed: seedInput.text
	property int clickX: 0
	property int clickY: 0

	width: 800
	height: 600

	color: "#030f14"

	Rectangle {
		id: background

		width: parent.width
		height: parent.height-50

		gradient: Gradient {
  		GradientStop { position: 0.0; color: "#061f29" }
    	GradientStop { position: 1.0; color: "#030f14" }
		}

		MouseArea {
  		id: area
  	  width: parent.width
  	  height: parent.height
			onClicked: {
				root.clickX = mouse.x
				root.clickY = mouse.y
			}
		}

		ParticleSystem {
			anchors.fill: parent

			ImageParticle {
				id: ip

				source: "image://star/FFFFFF"

				colorVariation: 0.2
				alpha: 0.6

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
			cache: false
			source: "image://karta/map.png"

			property int clicks: 0

			x: (parent.width - width)/2
			y: (parent.height - height)/2

			width: parent.width/1.6

			fillMode: Image.PreserveAspectFit
			clip: false

			MouseArea {
				anchors.fill: parent
				onClicked: {
					parent.clicks += 1
					parent.source = "image://karta/map" + karta.clicks
				}
			}
		}
	}

	Rectangle {
    id: form
		width: parent.width
		height: 50

		y: parent.height-50

		color: "#061f29"

    Row {
      id: row
			anchors.fill: parent
			anchors.margins: 10
			spacing: 20

			Text {
				id: seedText
				text: ctrl.message

				color: "#FFFFFF"
				font.pointSize: 12
				font.bold: true

				MouseArea {
					id: mouseArea
					anchors.fill: parent
					onReleased: ctrl.textReleased(seedInput)
				}
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
}

// End of the QML string`

// vim: ft=qml.go
