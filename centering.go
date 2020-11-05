package main

import (
	"image"

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
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e)
				drawLabels(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawLabels(gtx layout.Context, th *material.Theme) layout.Dimensions {
	gtx.Constraints.Min = image.Pt(0, 0)
	rec := op.Record(gtx.Ops) // Start recording  // HLcenter
	dimensions := material.H2(th, "I'm centered!").Layout(gtx)
	macro := rec.Stop() // End recording // HLcenter
	op.Affine(f32.Affine2D{}.Offset(f32.Point{
		X: float32(gtx.Constraints.Max.X-dimensions.Size.X) / 2,
		Y: float32(gtx.Constraints.Max.Y-dimensions.Size.Y) / 2,
	})).Add(gtx.Ops)
	macro.Add(gtx.Ops) // Replay operations // HLcenter
	return dimensions
}

// END OMIT
