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

// START OMIT
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

				lbl := text.Label{Face: faces.For(regular, ui.Sp(72)), Text: "Hello, World!"} // HLdraw
				lbl.Layout(ops, cs)                                                           // HLdraw

				w.Update(ops)
			}
		} // HLeventloop
	}()
	app.Main()
}

// END OMIT
