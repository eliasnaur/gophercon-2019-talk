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
		var faces measure.Faces
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, layout.RigidConstraints(e.Size))
				faces.Reset(gtx.Config)

				lbl := text.Label{Face: faces.For(regular, ui.Sp(72)), Text: "Hello, World!"} // HLdraw
				lbl.Layout(gtx)                                                               // HLdraw

				w.Update(gtx.Ops)
			}
		} // HLeventloop
	}()
	app.Main()
}

// END OMIT
