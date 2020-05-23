package main

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		var ops op.Ops
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e.Queue, e.Config, e.Size)
				drawRects(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		// Red.
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return drawRect(gtx, color.RGBA{A: 0xff, R: 0xff})
		}),

		// Green.
		layout.Flexed(0.25, func(gtx layout.Context) layout.Dimensions {
			return drawRect(gtx, color.RGBA{A: 0xff, G: 0xff})
		}),

		// Blue.
		layout.Flexed(0.25, func(gtx layout.Context) layout.Dimensions {
			return drawRect(gtx, color.RGBA{A: 0xff, B: 0xff})
		}),
	)
}

// END OMIT

func drawRect(gtx layout.Context, color color.RGBA) layout.Dimensions {
	cs := gtx.Constraints
	square := f32.Rectangle{
		Max: f32.Point{
			X: float32(cs.Max.X),
			Y: float32(cs.Max.Y),
		},
	}
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{Rect: square}.Add(gtx.Ops)
	return layout.Dimensions{Size: cs.Max}
}
