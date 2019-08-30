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

func drawRect(c ui.Config, ops *ui.Ops, color color.RGBA, cs layout.Constraints) layout.Dimensions {
	square := f32.Rectangle{
		Max: f32.Point{
			X: float32(cs.Width.Max),
			Y: float32(cs.Height.Max),
		},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
	return layout.Dimensions{Size: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}
}
