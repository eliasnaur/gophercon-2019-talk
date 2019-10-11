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
		for e := range w.Events() {
			if e, ok := e.(app.FrameEvent); ok {
				gtx.Reset(&e.Config, e.Size)
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
	th.H2("I'm centered!").Layout(gtx)
	dimensions := gtx.Dimensions
	macro.Stop() // End recording // HLcenter
	op.TransformOp{}.Offset(f32.Point{
		X: float32(gtx.Constraints.Width.Max-dimensions.Size.X) / 2,
		Y: float32(gtx.Constraints.Height.Max-dimensions.Size.Y) / 2,
	}).Add(gtx.Ops)
	macro.Add(gtx.Ops) // Replay operations // HLcenter
}

// END OMIT
