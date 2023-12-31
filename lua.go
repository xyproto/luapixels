package luapixels

import (
	"errors"
	"fmt"
	"time"

	lua "github.com/yuin/gopher-lua"
)

var blacklist []string

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

func quit(_ *lua.LState) int {
	shouldQuit = true
	return 0 // number of return values
}

// Lua binding for PlaySound
func playSound(L *lua.LState) int {
	frequency := L.ToNumber(1)
	duration := L.ToInt(2)
	PlaySound(float32(frequency), duration)
	return 0
}

func sleep(L *lua.LState) int {
	// Extract the correct number of nanoseconds
	duration := time.Duration(float64(L.ToNumber(1)) * 1000000000.0)
	// Wait and block the current thread of execution.
	time.Sleep(duration)
	return 0
}

// GetLuaGlobalString fetches the value of a global Lua variable as a string.
func GetLuaGlobalString(L *lua.LState, variableName string) (string, error) {
	global := L.GetGlobal(variableName)
	if global.Type() == lua.LTString { // success
		return global.String(), nil
	}
	return "", fmt.Errorf("global variable '%s' is not a string or doesn't exist", variableName)
}

// CallLuaFunction calls a Lua function by name.
func CallLuaFunction(L *lua.LState, funcName string) error {
	if hasS(blacklist, funcName) {
		return errors.New("no such function: " + funcName)
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(funcName),
		NRet:    0,
		Protect: true,
	}); err != nil {
		blacklist = append(blacklist, funcName)
		return err
	}
	return nil
}

// hasS checks if a slice of strings contains the given string
func hasS(xs []string, x string) bool {
	for _, e := range xs {
		if e == x {
			return true
		}
	}
	return false
}
