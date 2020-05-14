package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
)

// START QUEUE OMIT
func main() {
	go func() {
		w := app.NewWindow()
		button := new(Button)
		gtx := new(layout.Context)
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Queue, e.Config, e.Size)
				button.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// END QUEUE OMIT

type Button struct {
	pressed bool
}

// START OMIT
func (b *Button) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}

	col := color.RGBA{A: 0xff, R: 0xff}
	if b.pressed {
		col = color.RGBA{A: 0xff, G: 0xff}
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: 500, Y: 500}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	drawSquare(gtx.Ops, col)
}

// END OMIT

func drawSquare(ops *op.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
}
