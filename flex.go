package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				drawRects(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(gtx *layout.Context) {
	layout.Flex{}.Layout(gtx,
		// Red.
		layout.Flexed(0.5, func() {
			drawRect(gtx, color.RGBA{A: 0xff, R: 0xff})
		}),

		// Green.
		layout.Flexed(0.25, func() {
			drawRect(gtx, color.RGBA{A: 0xff, G: 0xff})
		}),

		// Blue.
		layout.Flexed(0.25, func() {
			drawRect(gtx, color.RGBA{A: 0xff, B: 0xff})
		}),
	)
}

// END OMIT

func drawRect(gtx *layout.Context, color color.RGBA) {
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
}
