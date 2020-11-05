package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	go func() {
		w := app.NewWindow()
		th := material.NewTheme(gofont.Collection())
		// START INIT OMIT
		editor := new(widget.Editor)
		editor.SetText("Hello, Gophers! Edit me.")
		// END INIT OMIT
		var ops op.Ops
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e) // HLdraw
				// START OMIT
				ed := material.Editor(th, editor, "Hint")
				ed.TextSize = unit.Sp(52)
				ed.Layout(gtx)
				// END OMIT
				e.Frame(gtx.Ops) // HLdraw
			}
		} // HLeventloop
	}()
	app.Main()
}
