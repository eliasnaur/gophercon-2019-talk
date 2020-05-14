package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow()
		gofont.Register()
		th := material.NewTheme()
		gtx := new(layout.Context)
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Queue, e.Config, e.Size)
				drawLabels(gtx, th) // HLdraw
				e.Frame(gtx.Ops)
			}
		}
		// END OMIT
	}()
	app.Main()
}

// START DRAW OMIT

func drawLabels(gtx *layout.Context, th *material.Theme) {
	material.H1(th, "One label").Layout(gtx)
	material.H1(th, "Another label").Layout(gtx)
}

// END DRAW OMIT
