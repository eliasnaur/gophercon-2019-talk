package main

import (
	"gioui.org/app"
	_ "gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

// START OMIT
func main() {
	go func() {
		w := app.NewWindow()
		th := material.NewTheme()
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		for e := range w.Events() {
			if e, ok := e.(app.FrameEvent); ok {
				gtx.Reset(&e.Config, e.Size)

				th.H1("Hello, World!").Layout(gtx) // HLdraw

				e.Frame(gtx.Ops)
			}
		} // HLeventloop
	}()
	app.Main()
}

// END OMIT
