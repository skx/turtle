// This is a simple turtle-application, allowing retro scripted graphics
// to be produced.
//
// We create a global image ("i") which we draw to, via the FORTH words
// "forward", "move", etc.
//
// The end result is output as a PNG.  However we also generate a GIF of
// the drawing process - updating at each step of the drawing process.
//
// NOTE: This is very slow for circles, for obvious reasons..
//

package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"os"

	"github.com/skx/foth/foth/eval"
)

// FORTH interpreter.
var e *eval.Eval

// Graphics-helper, for drawing into our image.
var g *Graphics

// Did an image get saved?
var saved bool

// Run the user-supplied script.
func runFORTH(eval *eval.Eval, path string) error {

	handle, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(handle)
	line, err := reader.ReadString(byte('\n'))
	for err == nil {

		// Evaluate
		err = e.Eval(line)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())

			// Reset our state, to allow recovery
			e.Reset()
		}

		// Repeat
		line, err = reader.ReadString(byte('\n'))
	}

	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	err = handle.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {

	// New PNG image - with white background
	i := image.NewRGBA(image.Rect(0, 0, 300, 300))
	c := color.RGBA{255, 255, 255, 255}
	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)

	// Create a new helper to draw into our image.
	//
	// Point is set to the middle of the image.
	g = NewGraphics(i, Position{150.0, 150.0})

	// Create a FORTH interpreter instance
	e = eval.New()

	// Give it access to our implementation, by defining new words
	e.Dictionary = append(e.Dictionary, eval.Word{Name: "direction", Function: direction})
	e.Dictionary = append(e.Dictionary, eval.Word{Name: "forward", Function: forward})
	e.Dictionary = append(e.Dictionary, eval.Word{Name: "move", Function: move})
	e.Dictionary = append(e.Dictionary, eval.Word{Name: "pen", Function: pen})
	e.Dictionary = append(e.Dictionary, eval.Word{Name: "save", Function: save})
	e.Dictionary = append(e.Dictionary, eval.Word{Name: "turn", Function: turn})

	// If we don't have any arguments abort.
	if len(os.Args) == 1 {
		fmt.Printf("Usage: turtle script1 script2 .. scriptN\n")
		return
	}

	// Process each user-defined script.
	for _, file := range os.Args[1:] {

		saved = false

		err := runFORTH(e, file)
		if err != nil {
			fmt.Printf("error running %s: %s\n", file, err.Error())
			return
		}

		if !saved {
			fmt.Printf("WARNING: Image not saved - did you forget to call save?")
		}
	}
}
