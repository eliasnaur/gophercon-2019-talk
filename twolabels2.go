package main

import (
	"gioui.org/ui"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/measure"
	"gioui.org/text"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/sfnt"
)

func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF)
		var faces measure.Faces
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, layout.RigidConstraints(e.Size))
				faces.Reset(gtx.Config)
				f := faces.For(regular, ui.Sp(122))
				drawLabels(gtx, f) // HLdraw
				w.Update(gtx.Ops)
			}
		}
		// END OMIT
	}()
	app.Main()
}

// START DRAW OMIT
func drawLabels(gtx *layout.Context, face text.Face) {
	gtx.Constraints.Height.Min = 0 // HLdraw
	lbl := text.Label{Face: face, Text: "One label"}
	lbl.Layout(gtx) // HLdraw
	dimensions := gtx.Dimensions
	ui.TransformOp{}.Offset(f32.Point{
		Y: float32(dimensions.Size.Y), // HLdraw
	}).Add(gtx.Ops)
	lbl2 := text.Label{Face: face, Text: "Another label"}
	lbl2.Layout(gtx)
}

// END DRAW OMIT
