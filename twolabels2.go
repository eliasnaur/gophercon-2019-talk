package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow()
		th := material.NewTheme(gofont.Collection())
		var ops op.Ops
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e)
				drawLabels(gtx, th) // HLdraw
				e.Frame(gtx.Ops)
			}
		}
		// END OMIT
	}()
	app.Main()
}

// START DRAW OMIT
func drawLabels(gtx layout.Context, th *material.Theme) {
	gtx.Constraints.Min.Y = 0                              // HLdraw
	dimensions := material.H1(th, "One label").Layout(gtx) // HLdraw
	op.Affine(f32.Affine2D{}.Offset(f32.Point{
		Y: float32(dimensions.Size.Y), // HLdraw
	})).Add(gtx.Ops)
	material.H1(th, "Another label").Layout(gtx)
}

// END DRAW OMIT
