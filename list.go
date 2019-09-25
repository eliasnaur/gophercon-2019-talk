package main

import (
	"fmt"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/layout"
	"gioui.org/ui/measure"
	"gioui.org/ui/text"

	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/sfnt"
)

func main() {
	go func() {
		w := app.NewWindow()
		regular, _ := sfnt.Parse(goregular.TTF)
		var faces measure.Faces
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
				gtx.Reset(&e.Config, layout.RigidConstraints(e.Size))
				faces.Reset(gtx.Config)
				f := faces.For(regular, ui.Sp(42))
				drawList(gtx, list, f)
				w.Update(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawList(gtx *layout.Context, list *layout.List, face text.Face) {
	const n = 1e6
	list.Layout(gtx, n, func(i int) {
		txt := fmt.Sprintf("List element #%d", i)

		lbl := text.Label{Face: face, Text: txt}
		lbl.Layout(gtx)
	})
}

// END OMIT
