// Implementation of words exported to FORTH.

package main

import (
	"math"
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
	err := g.WriteImage("turtle.png")
	if err != nil {
		return err
	}

	// Write the animation
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
