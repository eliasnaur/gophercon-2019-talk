package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		for e := range w.Events() {
			if e, ok := e.(app.FrameEvent); ok {
				gtx.Reset(&e.Config, e.Size)
				drawRects(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(gtx *layout.Context) {
	flex := layout.Flex{}

	red := flex.Flex(gtx, 0.5, func() {
		drawRect(gtx, color.RGBA{A: 0xff, R: 0xff})
	})

	green := flex.Flex(gtx, 0.25, func() {
		drawRect(gtx, color.RGBA{A: 0xff, G: 0xff})
	})

	blue := flex.Flex(gtx, 0.25, func() {
		drawRect(gtx, color.RGBA{A: 0xff, B: 0xff})
	})

	flex.Layout(gtx, red, green, blue)
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
