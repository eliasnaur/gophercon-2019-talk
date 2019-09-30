package main

import (
	"gioui.org/ui"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/measure"
	"gioui.org/text"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/sfnt"
)

func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF) // HLdraw
		// START INIT OMIT
		var faces measure.Faces // HLdraw
		editor := &text.Editor{
			Face: faces.For(regular, ui.Sp(52)),
		}
		editor.SetText("Hello, Gophercon! Edit me.")
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, layout.RigidConstraints(e.Size)) // HLdraw
				faces.Reset(gtx.Config)                               // HLdraw
				// START OMIT
				editor.Layout(gtx)
				// END OMIT
				w.Update(gtx.Ops) // HLdraw
			}
		} // HLeventloop
	}()
	app.Main()
}
