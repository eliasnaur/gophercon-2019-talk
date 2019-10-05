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

// START OMIT
func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF)
		var faces shape.Faces
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, e.Size)
				faces.Reset()

				lbl := text.Label{Face: faces.For(regular), Size: unit.Sp(72), Text: "Hello, World!"} // HLdraw
				lbl.Layout(gtx)                                                                       // HLdraw

				w.Update(gtx.Ops)
			}
		} // HLeventloop
	}()
	app.Main()
}

// END OMIT
