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
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, e.Size)
				faces.Reset()
				f := faces.For(regular)
				drawLabels(gtx, f, unit.Sp(72))
				w.Update(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawLabels(gtx *layout.Context, face text.Face, size unit.Value) {
	gtx.Constraints.Width.Min = 0
	gtx.Constraints.Height.Min = 0
	lbl := text.Label{Face: face, Size: size, Text: "I'm centered!"}
	var macro op.MacroOp  // HLcenter
	macro.Record(gtx.Ops) // Start recording  // HLcenter
	lbl.Layout(gtx)
	dimensions := gtx.Dimensions
	macro.Stop() // End recording // HLcenter
	op.TransformOp{}.Offset(f32.Point{
		X: float32(gtx.Constraints.Width.Max-dimensions.Size.X) / 2,
		Y: float32(gtx.Constraints.Height.Max-dimensions.Size.Y) / 2,
	}).Add(gtx.Ops)
	macro.Add(gtx.Ops) // Replay operations // HLcenter
}

// END OMIT
