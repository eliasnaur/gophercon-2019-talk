package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

// START OMIT
func main() {
	go func() {
		w := app.NewWindow()
		gofont.Register()
		th := material.NewTheme()
		var ops op.Ops
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e.Queue, e.Config, e.Size)

				material.H1(th, "Hello, World!").Layout(gtx) // HLdraw

				e.Frame(gtx.Ops)
			}
		} // HLeventloop
	}()
	app.Main()
}

// END OMIT
