package main

import (
	"gioui.org/app"
	"gioui.org/layout"
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
				gtx.Reset(&e.Config, layout.RigidConstraints(e.Size))
				faces.Reset(gtx.Config)
				f := faces.For(regular, unit.Sp(122))
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
	lbl := text.Label{Face: face, Text: "One label"}
	lbl.Layout(gtx)
	lbl2 := text.Label{Face: face, Text: "Another label"}
	lbl2.Layout(gtx)
}

// END DRAW OMIT
