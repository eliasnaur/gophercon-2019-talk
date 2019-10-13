package main

import (
	"fmt"

	"gioui.org/app"
	_ "gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow()
		th := material.NewTheme()
		// START INIT OMIT
		list := &layout.List{
			Axis: layout.Vertical,
		}
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(app.FrameEvent); ok {
				gtx.Reset(&e.Config, e.Size)
				drawList(gtx, list, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawList(gtx *layout.Context, list *layout.List, th *material.Theme) {
	const n = 1e6
	list.Layout(gtx, n, func(i int) {
		txt := fmt.Sprintf("List element #%d", i)

		th.H3(txt).Layout(gtx)
	})
}

// END OMIT
