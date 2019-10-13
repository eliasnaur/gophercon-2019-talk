package main

import (
	"gioui.org/app"
	_ "gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow()
		th := material.NewTheme()
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(app.FrameEvent); ok {
				gtx.Reset(&e.Config, e.Size)
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
	th.H1("One label").Layout(gtx)
	th.H1("Another label").Layout(gtx)
}

// END DRAW OMIT
