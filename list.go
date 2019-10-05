package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/text/shape"
	"gioui.org/unit"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/sfnt"
)

func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF)
		fml := &shape.Family{
			Regular: regular,
		}
		// START INIT OMIT
		list := &layout.List{
			Axis: layout.Vertical,
		}
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, e.Size)
				drawList(gtx, list, fml, unit.Sp(42))
				w.Update(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawList(gtx *layout.Context, list *layout.List, fml text.Family, size unit.Value) {
	const n = 1e6
	list.Layout(gtx, n, func(i int) {
		txt := fmt.Sprintf("List element #%d", i)

		lbl := text.Label{Size: size, Text: txt}
		lbl.Layout(gtx, fml)
	})
}

// END OMIT
