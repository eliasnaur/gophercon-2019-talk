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
		gofont.Register()
		th := material.NewTheme()
		gtx := new(layout.Context)
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Queue, e.Config, e.Size)
				drawLabels(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawLabels(gtx *layout.Context, th *material.Theme) {
	gtx.Constraints.Width.Min = 0
	gtx.Constraints.Height.Min = 0
	var macro op.MacroOp  // HLcenter
	macro.Record(gtx.Ops) // Start recording  // HLcenter
	material.H2(th, "I'm centered!").Layout(gtx)
	dimensions := gtx.Dimensions
	macro.Stop() // End recording // HLcenter
	op.TransformOp{}.Offset(f32.Point{
		X: float32(gtx.Constraints.Width.Max-dimensions.Size.X) / 2,
		Y: float32(gtx.Constraints.Height.Max-dimensions.Size.Y) / 2,
	}).Add(gtx.Ops)
	macro.Add() // Replay operations // HLcenter
}

// END OMIT
