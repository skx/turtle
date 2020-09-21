package main

import (
	"image"
	"image/color"
	"math"
)

// Graphics holds our state.
type Graphics struct {
	Image       *image.RGBA
	Pos         Position
	Orientation float64
	Color       color.Color
	Pen         bool
}

// Position holds the coordinates of our pen.
type Position struct {
	X, Y float64
}

// NewGraphics creates a new object.
func NewGraphics(i *image.RGBA, starting Position) (t *Graphics) {
	t = &Graphics{
		Image:       i,
		Pos:         starting,
		Orientation: 0.0,
		Color:       color.Black,
		Pen:         true,
	}

	return
}

// Move updates the location of the pen.
func (t *Graphics) Move(x float64, y float64) {
	t.Pos.X = x
	t.Pos.Y = y
}

// Forward moves the pen forwards, in a straight-line, from the
// current location - in the direction of travel.
func (t *Graphics) Forward(dist float64) {
	for i := 0; i < int(dist); i++ {
		if t.Pen {
			t.Image.Set(int(t.Pos.X), int(t.Pos.Y), t.Color)
		}

		x := 1.0 * math.Sin(t.Orientation)
		y := 1.0 * -math.Cos(t.Orientation)

		t.Pos = Position{t.Pos.X + x, t.Pos.Y + y}
	}
}

// Turn adds the specified number of degrees to our direction.
func (t *Graphics) Turn(radians float64) {
	t.Orientation += radians
}

// Direction sets the absolute direction of the pen.
func (t *Graphics) Direction(radians float64) {
	t.Orientation = radians
}

// PenUp lifts the pen, so movement will not draw anything.
func (t *Graphics) PenUp() {
	t.Pen = false
}

// PenDown lowers the pen, so movement will draw.
func (t *Graphics) PenDown() {
	t.Pen = true
}
