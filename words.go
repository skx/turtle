// Implementation of words exported to FORTH.

package main

import (
	"math"
)

// set the direction, absolutely.
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

	return nil
}

// pen teleports to x,y.
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

	return nil
}

// Set the pen up/down
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

// Save the image, and the animation.
func save() error {

	// write the image (PNG)
	err := g.WriteImage("turtle.png")
	if err != nil {
		return err
	}

	// Write the animation (GIF).
	err = g.WriteAnimation("turtle.gif")
	if err != nil {
		return err
	}

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

	return nil
}
