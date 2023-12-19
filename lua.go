package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// setPalette is a Go function to be called from Lua for setting a color in the palette.
func setPalette(L *lua.LState) int {
	index := L.ToInt(1)
	r := L.ToInt(2)
	g := L.ToInt(3)
	b := L.ToInt(4)
	//fmt.Printf("set color %d to (%d, %d, %d)\n", index, r, g, b)
	SetPaletteColor(index, r, g, b)
	return 0 // number of return values
}

// plotPixel is a Go function to be called from Lua for plotting a pixel on the screen.
func plotPixel(L *lua.LState) int {
	x := L.ToInt(1)
	y := L.ToInt(2)
	colorIndex := L.ToInt(3)
	//fmt.Printf("plot pixel color %d at (%d, %d)\n", colorIndex, x, y)
	PlotPixel(x, y, colorIndex)
	return 0 // number of return values
}

// drawBackground is a Go function to be called from Lua for setting the entire background to blue.
func drawBackground(L *lua.LState) int {
	r := L.ToInt(1)
	g := L.ToInt(2)
	b := L.ToInt(3)
	//fmt.Printf("draw background (%d, %d, %d)\n", r, g, b)
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
	if global.Type() == lua.LTString {
		return global.String(), nil
	}
	return "", fmt.Errorf("global variable '%s' is not a string or doesn't exist", variableName)
}

func quit(L *lua.LState) int {
	shouldQuit = true
	return 0
}

// InitLua initializes the Lua VM, registers Go functions, and loads the given Lua filename
func InitLua(luaFilename string) *lua.LState {
	L := lua.NewState()

	L.SetGlobal("set_pal", L.NewFunction(setPalette))
	L.SetGlobal("plot", L.NewFunction(plotPixel))
	L.SetGlobal("background", L.NewFunction(drawBackground))
	L.SetGlobal("quit", L.NewFunction(quit))

	if err := L.DoFile(luaFilename); err != nil {
		panic(err)
	}
	return L
}
