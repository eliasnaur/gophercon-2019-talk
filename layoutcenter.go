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

func init() {
	go func() {
		w := app.NewWindow(nil)
		regular, _ := sfnt.Parse(goregular.TTF)
		var cfg app.Config
		faces := &measure.Faces{Config: &cfg}
		ops := new(ui.Ops)
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				cfg = e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				f := faces.For(regular, ui.Sp(72))
				drawLabels(f, ops, cs)
				w.Draw(ops)
				faces.Frame()
			}
		}
	}()
}

func main() {
	app.Main()
}

// START OMIT
func drawLabels(face text.Face, ops *ui.Ops, cs layout.Constraints) {
	cs.Width.Min = 0
	cs.Height.Min = 0
	lbl := text.Label{Face: face, Text: "I'm centered!"}
	var macro ui.MacroOp // HLcenter
	macro.Record(ops)    // HLcenter
	dimensions := lbl.Layout(ops, cs)
	macro.Stop() // HLcenter
	ui.TransformOp{ui.Offset(f32.Point{
		X: float32(cs.Width.Max-dimensions.Size.X) / 2,
		Y: float32(cs.Height.Max-dimensions.Size.Y) / 2,
	})}.Add(ops)
	macro.Add(ops) // HLcenter
}

// END OMIT
