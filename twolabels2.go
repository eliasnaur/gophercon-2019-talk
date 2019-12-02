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
		gtx := layout.NewContext(w.Queue())
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				drawLabels(gtx, th) // HLdraw
				e.Frame(gtx.Ops)
			}
		}
		// END OMIT
	}()
	app.Main()
}

// START DRAW OMIT
func drawLabels(gtx *layout.Context, th *material.Theme) {
	gtx.Constraints.Height.Min = 0 // HLdraw
	th.H1("One label").Layout(gtx) // HLdraw
	dimensions := gtx.Dimensions
	op.TransformOp{}.Offset(f32.Point{
		Y: float32(dimensions.Size.Y), // HLdraw
	}).Add(gtx.Ops)
	th.H1("Another label").Layout(gtx)
}

// END DRAW OMIT
