package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := new(layout.Context)
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Queue, e.Config, e.Size)
				drawRects(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(gtx *layout.Context) {
	layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// Red.
		layout.Stacked(func() {
			drawRect(gtx, color.RGBA{A: 0xff, R: 0xff}, unit.Dp(50))
		}),

		// Green.
		layout.Stacked(func() {
			drawRect(gtx, color.RGBA{A: 0xff, G: 0xff}, unit.Dp(100))
		}),

		// Blue.
		layout.Stacked(func() {
			drawRect(gtx, color.RGBA{A: 0xff, B: 0xff}, unit.Dp(150))
		}),
	)
}

// END OMIT

func drawRect(gtx *layout.Context, color color.RGBA, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		cs := gtx.Constraints
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(cs.Width.Max),
				Y: float32(cs.Height.Max),
			},
		}
		paint.ColorOp{Color: color}.Add(gtx.Ops)
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}
	})
}
