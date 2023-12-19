package main

import (
	"image/color"

	"github.com/fzipp/vga"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	scale        = 4
	width        = 320
	windowWidth  = width * scale
	height       = 200
	windowHeight = height * scale
	paletteSize  = 256
)

var (
	palette []color.Color
)

// Initialize GLFW and OpenGL
func initGraphics(windowTitle string) *glfw.Window {
	palette = vga.DefaultPalette

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
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
	palette[index] = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

// GetPaletteColor retrieves the r,g,b colors of a given palette index.
func GetPaletteColor(index int) (int, int, int) {
	color := palette[index].(color.RGBA)
	return int(color.R), int(color.G), int(color.B)
}

// PlotPixel places a pixel with the given colorIndex
func PlotPixel(x, y, colorIndex int) {
	colorRGBA := palette[colorIndex].(color.RGBA)

	// Calculate the coordinates for the quad's vertices
	left := float32(x)                 //- 0.1
	right := left + float32(scale)     //- 0.1
	top := float32(y) - float32(scale) //- 0.1
	bottom := top + float32(scale)     //- 0.1

	// Set the OpenGL color
	gl.Color3ub(uint8(colorRGBA.R), uint8(colorRGBA.G), uint8(colorRGBA.B))

	// Draw a quad at the specified position
	gl.Begin(gl.QUADS)
	gl.Vertex2f(left, top)
	gl.Vertex2f(right, top)
	gl.Vertex2f(right, bottom)
	gl.Vertex2f(left, bottom)
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
	gl.ClearColor(float32(r)/255.0, float32(g)/255.0, float32(b)/255.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
