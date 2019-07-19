package main

import (
	"gioui.org/ui"
	"gioui.org/ui/app"
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
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				cfg = e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				f := faces.For(regular, ui.Sp(122))
				drawLabels(f, ops, cs) // HLdraw
				w.Draw(ops)
				faces.Frame()
			}
		}
		// END OMIT
	}()
}

func main() {
	app.Main()
}

// START DRAW OMIT
func drawLabels(face text.Face, ops *ui.Ops, cs layout.Constraints) {
	lbl := text.Label{Face: face, Text: "One label"}
	lbl.Layout(ops, cs)
	lbl2 := text.Label{Face: face, Text: "Another label"}
	lbl2.Layout(ops, cs)
}

// END DRAW OMIT
