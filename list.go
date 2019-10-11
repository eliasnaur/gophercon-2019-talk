package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/text/opentype"
	"gioui.org/widget/material"

	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	go func() {
		w := app.NewWindow()
		shaper := new(text.Shaper)
		shaper.Register(text.Font{}, opentype.Must(
			opentype.Parse(goregular.TTF),
		))
		th := material.NewTheme(shaper)
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
