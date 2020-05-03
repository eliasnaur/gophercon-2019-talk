package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	go func() {
		w := app.NewWindow()
		gofont.Register()
		th := material.NewTheme()
		// START INIT OMIT
		editor := new(widget.Editor)
		editor.SetText("Hello, Gophers! Edit me.")
		// END INIT OMIT
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size) // HLdraw
				// START OMIT
				ed := material.Editor(th, "Hint")
				ed.TextSize = unit.Sp(52)
				ed.Layout(gtx, editor)
				// END OMIT
				e.Frame(gtx.Ops) // HLdraw
			}
		} // HLeventloop
	}()
	app.Main()
}
