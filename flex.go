package main

import (
	"image"
	"image/color"

	"gioui.org/ui/app"
	"gioui.org/ui/f32"
	"gioui.org/ui/layout"
	"gioui.org/ui/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				gtx.Reset(&e.Config, layout.RigidConstraints(e.Size))
				drawRects(gtx)
				w.Update(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(gtx *layout.Context) {
	flex := layout.Flex{}
	flex.Init(gtx)

	red := flex.Flexible(0.5, func() {
		drawRect(gtx, color.RGBA{A: 0xff, R: 0xff})
	})

	green := flex.Flexible(0.25, func() {
		drawRect(gtx, color.RGBA{A: 0xff, G: 0xff})
	})

	blue := flex.Flexible(0.25, func() {
		drawRect(gtx, color.RGBA{A: 0xff, B: 0xff})
	})

	flex.Layout(red, green, blue)
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
