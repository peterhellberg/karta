package main

import (
	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/gl/2.1"
)

func init() {
	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(r *GoRect, obj qml.Object) {
			r.Object = obj
		},
	}})
}

type GoRect struct {
	qml.Object
}

func (r *GoRect) Paint(p *qml.Painter) {
	gl := GL.API(p)

	width := float32(r.Int("width"))
	height := float32(r.Int("height"))

	gl.Enable(GL.BLEND)
	gl.BlendFunc(GL.SRC_ALPHA, GL.ONE_MINUS_SRC_ALPHA)
	gl.Color4ub(0xff, 0x66, 0x00, 0xaa)
	gl.Begin(GL.QUADS)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(0, height)
	gl.End()

	gl.LineWidth(2.5)
	gl.Color4f(0.0, 0.0, 0.0, 1.0)
	gl.Color4ub(0xff, 0x66, 0x00, 0xff)
	gl.Begin(GL.LINES)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(0, height)
	gl.End()
}
