package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := app.NewWindow()
		var ops op.Ops
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e)
				drawRects(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(gtx layout.Context) layout.Dimensions {
	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// Red.
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return drawRect(gtx, color.RGBA{A: 0xff, R: 0xff}, unit.Dp(50))
		}),

		// Green.
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return drawRect(gtx, color.RGBA{A: 0xff, G: 0xff}, unit.Dp(100))
		}),

		// Blue.
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return drawRect(gtx, color.RGBA{A: 0xff, B: 0xff}, unit.Dp(150))
		}),
	)
}

// END OMIT

func drawRect(gtx layout.Context, color color.RGBA, inset unit.Value) layout.Dimensions {
	in := layout.UniformInset(inset)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		cs := gtx.Constraints
		square := image.Rectangle{Max: cs.Max}
		paint.FillShape(gtx.Ops, color, clip.Rect(square).Op())
		return layout.Dimensions{Size: cs.Max}
	})
}
