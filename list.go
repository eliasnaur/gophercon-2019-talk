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

func init() {
	go func() {
		w := app.NewWindow(nil)
		regular, _ := sfnt.Parse(goregular.TTF)
		var cfg app.Config
		faces := &measure.Faces{Config: &cfg}
		ops := new(ui.Ops)
		list := &layout.List{
			Axis:   layout.Vertical,
			Config: &cfg,
			Inputs: w.Queue(),
		}
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				cfg = e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				f := faces.For(regular, ui.Sp(42))
				drawList(list, f, ops, cs)
				w.Draw(ops)
				faces.Frame()
			}
		}
	}()
}

func main() {
	app.Main()
}

// START OMIT
func drawList(list *layout.List, face text.Face, ops *ui.Ops, cs layout.Constraints) {
	const n = 1e6
	for list.Init(ops, cs, n); list.More(); list.Next() { // HLlist
		txt := fmt.Sprintf("List element #%d", list.Index()) // HLlist
		lbl := text.Label{Face: face, Text: txt}
		dims := lbl.Layout(ops, list.Constraints()) // HLlist
		list.Elem(dims)                             // HLlist
	}
	list.Layout() // HLlist
}

// END OMIT
