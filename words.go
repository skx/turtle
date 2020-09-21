// Implementation of words exported to FORTH.

package main

import (
	"image/gif"
	"image/png"
	"math"
	"os"
)

// set the direction
//
// NOTE: This doesn't have any bearing on the animation.
func direction() error {
	val, err := e.Stack.Pop()
	if err != nil {
		return err
	}

	// degree -> radians
	val = val * (math.Pi / 180)
	g.Direction(val)

	return nil
}

// move forwards.
func forward() error {
	val, err := e.Stack.Pop()
	if err != nil {
		return err
	}
	g.Forward(val)

	// new image for GIF
	appendAnimation()

	return nil
}

// pen moves to x,y.
func move() error {
	y, err := e.Stack.Pop()
	if err != nil {
		return err
	}
	x, err := e.Stack.Pop()
	if err != nil {
		return err
	}
	g.Move(x, y)

	// new image for GIF
	appendAnimation()

	return nil
}

// Set the pen up/down
//
// NOTE: This doesn't have any bearing on the animation.
func pen() error {
	val, err := e.Stack.Pop()
	if err != nil {
		return err
	}
	if val == 0 {
		g.PenUp()
	} else {
		g.PenDown()
	}
	return nil
}

// save the image
func save() error {

	// write the PNG
	f, err := os.Create("turtle.png")
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, g.Image)

	// write the gif
	h, err2 := os.OpenFile("turtle.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err2 != nil {
		return err2
	}
	defer h.Close()
	gif.EncodeAll(h, outGif)

	saved = true

	return nil
}

// turn
func turn() error {
	val, err := e.Stack.Pop()
	if err != nil {
		return err
	}

	// degree -> radians
	val = val * (math.Pi / 180)
	g.Turn(val)

	// new image for GIF
	appendAnimation()

	return nil
}
