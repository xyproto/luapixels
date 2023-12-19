package luapixels

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// setPalette is a Go function to be called from Lua for setting a color in the palette.
func setPalette(L *lua.LState) int {
	colorIndex := byte(L.ToInt(1))
	r := byte(L.ToInt(2))
	g := byte(L.ToInt(3))
	b := byte(L.ToInt(4))
	SetPaletteColor(colorIndex, r, g, b)
	return 0 // number of return values
}

// putPixel is a Go function to be called from Lua for plotting a pixel
func putPixel(L *lua.LState) int {
	x := L.ToInt(1)
	y := L.ToInt(2)
	colorIndex := byte(L.ToInt(3))
	PutPixel(x, y, colorIndex)
	return 0 // number of return values
}

// getPixel is a Go function to be called from Lua for getting the color of a pixel
func getPixel(L *lua.LState) int {
	x := L.ToInt(1)
	y := L.ToInt(2)
	L.Push(lua.LNumber(GetPixel(x, y)))
	return 1 // number of return values
}

// drawBackground is a Go function to be called from Lua for setting the entire background to blue.
func drawBackground(L *lua.LState) int {
	r := byte(L.ToInt(1))
	g := byte(L.ToInt(2))
	b := byte(L.ToInt(3))
	DrawBackground(r, g, b)
	return 0 // number of return values
}

// CallLuaFunction calls a Lua function by name.
func CallLuaFunction(L *lua.LState, functionName string) {
	L.CallByParam(lua.P{
		Fn:      L.GetGlobal(functionName),
		NRet:    0,
		Protect: true,
	})
}

// GetLuaGlobalString fetches the value of a global Lua variable as a string.
func GetLuaGlobalString(L *lua.LState, variableName string) (string, error) {
	global := L.GetGlobal(variableName)
	if global.Type() == lua.LTString { // success
		return global.String(), nil
	}
	return "", fmt.Errorf("global variable '%s' is not a string or doesn't exist", variableName)
}

func quit(_ *lua.LState) int {
	shouldQuit = true
	return 0 // number of return values
}
