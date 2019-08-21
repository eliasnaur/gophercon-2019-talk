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
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				cfg = &e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				faces.Reset(cfg)
				f := faces.For(regular, ui.Sp(72))
				drawLabels(f, ops, cs)
				w.Update(ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawLabels(face text.Face, ops *ui.Ops, cs layout.Constraints) {
	cs.Width.Min = 0
	cs.Height.Min = 0
	lbl := text.Label{Face: face, Text: "I'm centered!"}
	var macro ui.MacroOp // HLcenter
	macro.Record(ops)    // Start recording  // HLcenter
	dimensions := lbl.Layout(ops, cs)
	macro.Stop() // End recording // HLcenter
	ui.TransformOp{}.Offset(f32.Point{
		X: float32(cs.Width.Max-dimensions.Size.X) / 2,
		Y: float32(cs.Height.Max-dimensions.Size.Y) / 2,
	}).Add(ops)
	macro.Add(ops) // Replay operations // HLcenter
}

// END OMIT
