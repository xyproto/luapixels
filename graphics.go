package luapixels

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

type PixelSlice [width * height]byte

var (
	palette []color.Color
	pixels  PixelSlice
)

func init() {
	palette = vga.DefaultPalette
}

func InitGL(window *glfw.Window) error {
	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		return err
	}

	// Set initial viewport and projection
	updateViewportAndProjection(window)

	return nil
}

func updateViewportAndProjection(window *glfw.Window) {
	winWidth, winHeight := window.GetSize()
	// Calculate the new viewport to keep the content centered
	viewportWidth := scale * width
	viewportHeight := scale * height
	viewportX := (winWidth - viewportWidth) / 2
	viewportY := (winHeight - viewportHeight) / 2

	gl.Viewport(int32(viewportX), int32(viewportY), int32(viewportWidth), int32(viewportHeight))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, width, height, 0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

// SetPaletteColor sets a color in the palette.
func SetPaletteColor(colorIndex, r, g, b byte) {
	palette[colorIndex] = color.RGBA{r, g, b, 0xff}
}

// GetPaletteColor retrieves the r,g,b colors of a given palette index.
func GetPaletteColor(colorIndex byte) (byte, byte, byte) {
	color := palette[colorIndex].(color.RGBA)
	return byte(color.R), byte(color.G), byte(color.B)
}

// glSetColor will set the current OpenGL color
func GLSetColor(c color.RGBA) {
	gl.Color3ub(byte(c.R), byte(c.G), byte(c.B))
}

// PutPixel places a pixel with the given colorIndex
func PutPixel(x, y int, colorIndex byte) {
	colorRGBA := palette[colorIndex].(color.RGBA)

	// Calculate the coordinates for the quad's vertices
	left := float32(x)
	right := left + float32(scale)
	top := float32(y) - float32(scale)
	bottom := top + float32(scale)

	// Set the OpenGL color
	GLSetColor(colorRGBA)

	// Draw a quad at the specified position
	gl.Begin(gl.QUADS)
	gl.Vertex2f(left, top)
	gl.Vertex2f(right, top)
	gl.Vertex2f(right, bottom)
	gl.Vertex2f(left, bottom)
	gl.End()

	pixels[width*y+x] = colorIndex
}

// GetPixel returns the color index of the given (x,y) position
func GetPixel(x, y int) byte {
	return pixels[width*y+x]
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
func DrawBackground(r, g, b byte) {
	gl.ClearColor(float32(r)/255.0, float32(g)/255.0, float32(b)/255.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
