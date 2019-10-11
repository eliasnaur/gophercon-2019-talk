package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/text/opentype"
	"gioui.org/widget/material"

	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	go func() {
		w := app.NewWindow()
		shaper := new(text.Shaper)
		shaper.Register(text.Font{}, opentype.Must(
			opentype.Parse(goregular.TTF),
		))
		th := material.NewTheme(shaper)
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(app.FrameEvent); ok {
				gtx.Reset(&e.Config, e.Size)
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
