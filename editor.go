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
		regular, _ := sfnt.Parse(goregular.TTF) // HLdraw
		// START INIT OMIT
		var faces shape.Faces // HLdraw
		editor := &text.Editor{
			Face: faces.For(regular),
			Size: unit.Sp(52),
		}
		editor.SetText("Hello, Gophercon! Edit me.")
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, e.Size) // HLdraw
				faces.Reset()                // HLdraw
				// START OMIT
				editor.Layout(gtx)
				// END OMIT
				w.Update(gtx.Ops) // HLdraw
			}
		} // HLeventloop
	}()
	app.Main()
}
