package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	lua "github.com/yuin/gopher-lua"
)

const luaFilename = "index.lua"

var shouldQuit = false

func mainProgram() int {
	runtime.LockOSThread()

	L := InitLua(luaFilename)
	defer L.Close()

	windowTitle, err := GetLuaGlobalString(L, "window_title")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s must declare a top level windowTitle variable\n", luaFilename)
		return 1
	}

	window := initGraphics(windowTitle)
	defer glfw.Terminate()

	CallLuaFunction(L, "at_start")

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			L.SetGlobal("last_key", lua.LNumber(key))
			CallLuaFunction(L, "at_keypress")
		}
	})

	for !window.ShouldClose() && !shouldQuit {
		ClearScreen()

		CallLuaFunction(L, "at_every_frame")

		UpdateScreen(window)
		glfw.PollEvents()
	}

	CallLuaFunction(L, "at_end")

	return 0
}

func main() {
	if retVal := mainProgram(); retVal != 0 {
		os.Exit(retVal)
	}
}
