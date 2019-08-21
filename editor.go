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

func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF) // HLdraw
		var cfg ui.Config                       // HLdraw
		ops := new(ui.Ops)                      // HLdraw
		// START INIT OMIT
		var faces measure.Faces // HLdraw
		editor := &text.Editor{
			Face: faces.For(regular, ui.Sp(52)),
		}
		editor.SetText("Hello, Gophercon! Edit me.")
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				cfg = &e.Config                       // HLdraw
				cs := layout.RigidConstraints(e.Size) // HLdraw
				ops.Reset()                           // HLdraw
				faces.Reset(cfg)                      // HLdraw
				queue := w.Queue()
				// START OMIT
				editor.Layout(cfg, queue, ops, cs)
				// END OMIT
				w.Update(ops) // HLdraw
			}
		} // HLeventloop
	}()
	app.Main()
}
