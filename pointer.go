package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/op"
	"gioui.org/op/paint"
)

// START QUEUE OMIT
func main() {
	go func() {
		w := app.NewWindow()
		button := new(Button)
		ops := new(op.Ops) // HLops
		for e := range w.Events() {
			if _, ok := e.(app.UpdateEvent); ok {
				ops.Reset()
				queue := w.Queue() // HLqueue
				button.Layout(queue, ops)
				w.Update(ops)
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
func (b *Button) Layout(queue event.Queue, ops *op.Ops) {
	for _, e := range queue.Events(b) { // HLevent
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
	pointer.RectAreaOp{ // HLevent
		Rect: image.Rectangle{Max: image.Point{X: 500, Y: 500}}, // HLevent
	}.Add(ops) // HLevent
	pointer.InputOp{Key: b}.Add(ops) // HLevent
	drawSquare(ops, col)
}

// END OMIT

func drawSquare(ops *op.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
}
