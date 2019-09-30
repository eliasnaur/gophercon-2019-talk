package main

import (
	"image"
	"image/color"

	"gioui.org/ui"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/paint"
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
	stack := layout.Stack{Alignment: layout.Center}
	stack.Init(gtx)

	red := stack.Rigid(func() {
		drawRect(gtx, color.RGBA{A: 0xff, R: 0xff}, ui.Dp(50))
	})

	green := stack.Rigid(func() {
		drawRect(gtx, color.RGBA{A: 0xff, G: 0xff}, ui.Dp(100))
	})

	blue := stack.Rigid(func() {
		drawRect(gtx, color.RGBA{A: 0xff, B: 0xff}, ui.Dp(150))
	})

	stack.Layout(red, green, blue)
}

// END OMIT

func drawRect(gtx *layout.Context, color color.RGBA, inset ui.Value) {
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
