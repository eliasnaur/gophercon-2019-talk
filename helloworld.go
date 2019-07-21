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
func init() { // HLinitfunc
	go func() { // HLinitfunc
		w := app.NewWindow(nil)
		regular, _ := sfnt.Parse(goregular.TTF) // HLdraw
		var cfg ui.Config                       // HLdraw
		var faces measure.Faces                 // HLdraw
		ops := new(ui.Ops)                      // HLdraw
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				cfg = &e.Config                                                               // HLdraw
				cs := layout.RigidConstraints(e.Size)                                         // HLdraw
				ops.Reset()                                                                   // HLdraw
				faces.Reset(cfg)                                                              // HLdraw
				lbl := text.Label{Face: faces.For(regular, ui.Sp(72)), Text: "Hello, World!"} // HLdraw
				lbl.Layout(ops, cs)                                                           // HLdraw
				w.Draw(ops)                                                                   // HLdraw
			}
		} // HLeventloop
	}() // HLinitfunc
} // HLinitfunc

func main() { // HLmain
	app.Main() // HLmain
} // HLmain

// END OMIT
