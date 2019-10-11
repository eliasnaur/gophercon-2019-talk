package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/text/opentype"
	"gioui.org/widget/material"
	"golang.org/x/image/font/gofont/goregular"
)

// START OMIT
func main() {
	go func() {
		w := app.NewWindow()
		shaper := new(text.Shaper)
		shaper.Register(text.Font{}, opentype.Must(
			opentype.Parse(goregular.TTF),
		))
		th := material.NewTheme(shaper)
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
