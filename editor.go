package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	_ "gioui.org/font/gofont"
)

func main() {
	go func() {
		w := app.NewWindow()
		th := material.NewTheme()
		// START INIT OMIT
		editor := new(widget.Editor)
		editor.SetText("Hello, Gophercon! Edit me.")
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(app.FrameEvent); ok {
				gtx.Reset(&e.Config, e.Size) // HLdraw
				// START OMIT
				ed := th.Editor("Hint")
				ed.Font.Size = unit.Sp(52)
				ed.Layout(gtx, editor)
				// END OMIT
				e.Frame(gtx.Ops) // HLdraw
			}
		} // HLeventloop
	}()
	app.Main()
}
