package main

import (
	"image"
	"image/color"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/paint"
	"gioui.org/ui/f32"
	"gioui.org/ui/layout"
)

func main() {
	go func() {
		w := app.NewWindow(nil)
		var cfg app.Config
		ops := new(ui.Ops)
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				cfg = e.Config
				cs := layout.RigidConstraints(e.Size)
				ops.Reset()
				drawRects(&cfg, ops, cs)
				w.Draw(ops)
			}
		}
	}()
	app.Main()
}

// START OMIT
func drawRects(c ui.Config, ops *ui.Ops, cs layout.Constraints) {
	flex := layout.Flex{}
	flex.Init(ops, cs)

	cs = flex.Flexible(0.5)
	dimensions := drawRect(c, ops, color.RGBA{A: 0xff, R: 0xff}, cs)
	red := flex.End(dimensions)

	cs = flex.Flexible(0.25)
	dimensions = drawRect(c, ops, color.RGBA{A: 0xff, G: 0xff}, cs)
	green := flex.End(dimensions)

	cs = flex.Flexible(0.25)
	dimensions = drawRect(c, ops, color.RGBA{A: 0xff, B: 0xff}, cs)
	blue := flex.End(dimensions)

	flex.Layout(red, green, blue)
}

// END OMIT

func drawRect(c ui.Config, ops *ui.Ops, color color.RGBA, cs layout.Constraints) layout.Dimens {
	square := f32.Rectangle{
		Max: f32.Point{
			X: float32(cs.Width.Max),
			Y: float32(cs.Height.Max),
		},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
	return layout.Dimens{Size: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}
}
