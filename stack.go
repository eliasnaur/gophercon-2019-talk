package main

import (
	"image"
	"image/color"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/f32"
	"gioui.org/ui/layout"
	"gioui.org/ui/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		var cfg app.Config
		ops := new(ui.Ops)
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				cfg = e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				drawRects(&cfg, ops, cs)
				w.Update(ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(c ui.Config, ops *ui.Ops, cs layout.Constraints) {
	stack := layout.Stack{Alignment: layout.Center}
	stack.Init(ops, cs)

	cs = stack.Rigid()
	dimensions := drawRect(c, ops, color.RGBA{A: 0xff, R: 0xff}, ui.Dp(50), cs)
	red := stack.End(dimensions)

	cs = stack.Rigid()
	dimensions = drawRect(c, ops, color.RGBA{A: 0xff, G: 0xff}, ui.Dp(100), cs)
	green := stack.End(dimensions)

	cs = stack.Rigid()
	dimensions = drawRect(c, ops, color.RGBA{A: 0xff, B: 0xff}, ui.Dp(150), cs)
	blue := stack.End(dimensions)

	stack.Layout(red, green, blue)
}

// END OMIT

func drawRect(c ui.Config, ops *ui.Ops, color color.RGBA, inset ui.Value, cs layout.Constraints) layout.Dimensions {
	in := layout.UniformInset(inset)
	cs = in.Begin(c, ops, cs)
	square := f32.Rectangle{
		Max: f32.Point{
			X: float32(cs.Width.Max),
			Y: float32(cs.Height.Max),
		},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
	dimens := layout.Dimensions{Size: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}
	dimens = in.End(dimens)
	return dimens
}
