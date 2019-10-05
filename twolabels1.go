package main

import (
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
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		// START OMIT
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, e.Size)
				drawLabels(gtx, fml, unit.Sp(122)) // HLdraw
				w.Update(gtx.Ops)
			}
		}
		// END OMIT
	}()
	app.Main()
}

// START DRAW OMIT
func drawLabels(gtx *layout.Context, fml text.Family, size unit.Value) {
	lbl := text.Label{Size: size, Text: "One label"}
	lbl.Layout(gtx, fml)
	lbl2 := text.Label{Text: "Another label"}
	lbl2.Layout(gtx, fml)
}

// END DRAW OMIT
