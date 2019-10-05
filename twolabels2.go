package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/text/shape"
	"gioui.org/unit"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/sfnt"
)

func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF)
		var faces shape.Faces
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, e.Size)
				faces.Reset()
				f := faces.For(regular)
				drawLabels(gtx, f, unit.Sp(122)) // HLdraw
				w.Update(gtx.Ops)
			}
		}
		// END OMIT
	}()
	app.Main()
}

// START DRAW OMIT
func drawLabels(gtx *layout.Context, face text.Face, size unit.Value) {
	gtx.Constraints.Height.Min = 0 // HLdraw
	lbl := text.Label{Face: face, Size: size, Text: "One label"}
	lbl.Layout(gtx) // HLdraw
	dimensions := gtx.Dimensions
	op.TransformOp{}.Offset(f32.Point{
		Y: float32(dimensions.Size.Y), // HLdraw
	}).Add(gtx.Ops)
	lbl2 := text.Label{Face: face, Text: "Another label"}
	lbl2.Layout(gtx)
}

// END DRAW OMIT
