package main

import (
	"fmt"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/input"
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
		var cfg ui.Config
		var faces measure.Faces
		ops := new(ui.Ops)
		// START INIT OMIT
		list := &layout.List{
			Axis: layout.Vertical,
		}
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				cfg = &e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				faces.Reset(cfg)
				f := faces.For(regular, ui.Sp(42))
				drawList(cfg, w.Queue(), list, f, ops, cs)
				w.Update(ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawList(c ui.Config, q input.Queue, list *layout.List, face text.Face, ops *ui.Ops, cs layout.Constraints) {
	const n = 1e6
	for list.Init(c, q, ops, cs, n); list.More(); list.Next() {
		txt := fmt.Sprintf("List element #%d", list.Index())

		lbl := text.Label{Face: face, Text: txt}
		dims := lbl.Layout(ops, list.Constraints())

		list.End(dims)
	}
	list.Layout()
}

// END OMIT
