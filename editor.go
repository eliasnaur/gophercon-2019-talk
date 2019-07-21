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
		editor := &text.Editor{
			Face: faces.For(regular, ui.Sp(22)),
		}
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				cfg = &e.Config                       // HLdraw
				cs := layout.RigidConstraints(e.Size) // HLdraw
				ops.Reset()                           // HLdraw
				faces.Reset(cfg)                      // HLdraw
				editor.Layout(cfg, w.Queue(), ops, cs)
				w.Draw(ops) // HLdraw
			}
		} // HLeventloop
	}() // HLinitfunc
} // HLinitfunc

func main() { // HLmain
	app.Main() // HLmain
} // HLmain

// END OMIT
