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
		gofont.Register()
		th := material.NewTheme()
		var ops op.Ops
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e.Queue, e.Config, e.Size)
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
	var macro op.MacroOp  // HLcenter
	macro.Record(gtx.Ops) // Start recording  // HLcenter
	dimensions := material.H2(th, "I'm centered!").Layout(gtx)
	macro.Stop() // End recording // HLcenter
	op.TransformOp{}.Offset(f32.Point{
		X: float32(gtx.Constraints.Max.X-dimensions.Size.X) / 2,
		Y: float32(gtx.Constraints.Max.Y-dimensions.Size.Y) / 2,
	}).Add(gtx.Ops)
	macro.Add() // Replay operations // HLcenter
	return dimensions
}

// END OMIT
