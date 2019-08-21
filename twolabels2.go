package main

import (
	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/f32"
	"gioui.org/ui/layout"
	"gioui.org/ui/measure"
	"gioui.org/ui/text"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/sfnt"
)

func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF)
		var cfg ui.Config
		var faces measure.Faces
		ops := new(ui.Ops)
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				cfg = &e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				faces.Reset(cfg)
				f := faces.For(regular, ui.Sp(122))
				drawLabels(f, ops, cs) // HLdraw
				w.Update(ops)
			}
		}
		// END OMIT
	}()
	app.Main()
}

// START DRAW OMIT
func drawLabels(face text.Face, ops *ui.Ops, cs layout.Constraints) {
	cs.Height.Min = 0 // HLdraw
	lbl := text.Label{Face: face, Text: "One label"}
	dimensions := lbl.Layout(ops, cs) // HLdraw
	ui.TransformOp{}.Offset(f32.Point{
		Y: float32(dimensions.Size.Y), // HLdraw
	}).Add(ops)
	lbl2 := text.Label{Face: face, Text: "Another label"}
	lbl2.Layout(ops, cs)
}

// END DRAW OMIT
