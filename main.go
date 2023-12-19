package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func mainProgram() int {
	// TODO: Is this needed with glfw?
	runtime.LockOSThread()

	const luaFilename = "index.lua"

	L := InitLua(luaFilename)
	defer L.Close()

	windowTitle, err := GetLuaGlobalString(L, "windowTitle")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s must declare a windowTitle variable\n", luaFilename)
		return 1
	}

	window := initGraphics(windowTitle)
	defer glfw.Terminate()

	CallLuaFunction(L, "atStart")

	for !window.ShouldClose() {
		ClearScreen()

		CallLuaFunction(L, "everyFrame")

		UpdateScreen(window)
		glfw.PollEvents()
	}

	CallLuaFunction(L, "atEnd")

	return 0
}

func main() {
	if retVal := mainProgram(); retVal != 0 {
		os.Exit(retVal)
	}
}
