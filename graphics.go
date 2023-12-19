package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 320
	height = 200
)

var (
	palette [256][3]float32 // Palette of 256 colors
)

// Initialize GLFW and OpenGL
func initGraphics(windowTitle string) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(width, height, windowTitle, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Set up orthographic projection
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, width, height, 0, -1, 1) // Set coordinate system
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	return window
}

// SetPaletteColor sets a color in the palette.
func SetPaletteColor(index int, r, g, b int) {
	palette[index] = [3]float32{float32(r) / 255, float32(g) / 255, float32(b) / 255}
}

// PlotPixel plots a pixel on the screen using a color from the palette.
func PlotPixel(x, y, colorIndex int) {
	gl.PointSize(5.0)
	gl.Color3fv(&palette[colorIndex][0])
	gl.Begin(gl.POINTS)
	gl.Vertex2f(float32(x), float32(y))
	gl.End()
}

// UpdateScreen swaps the buffers and displays the rendered frame.
func UpdateScreen(window *glfw.Window) {
	window.SwapBuffers()
}

// ClearScreen clears the window.
func ClearScreen() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// DrawBackground sets the entire background to a specified color
func DrawBackground(r, g, b int) {
	// Normalize color components to be between 0 and 1
	gl.ClearColor(float32(r)/255.0, float32(g)/255.0, float32(b)/255.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
