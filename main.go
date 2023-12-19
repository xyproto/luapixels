package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const luaFilename = "index.lua"

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

	for !window.ShouldClose() {
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
