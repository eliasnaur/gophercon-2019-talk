package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow()
		th := material.NewTheme(gofont.Collection())
		// START INIT OMIT
		list := &layout.List{
			Axis: layout.Vertical,
		}
		var ops op.Ops
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e)
				drawList(gtx, list, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawList(gtx layout.Context, list *layout.List, th *material.Theme) layout.Dimensions {
	const n = 1e6
	return list.Layout(gtx, n, func(gtx layout.Context, i int) layout.Dimensions {
		txt := fmt.Sprintf("List element #%d", i)

		return material.H3(th, txt).Layout(gtx)
	})
}

// END OMIT
