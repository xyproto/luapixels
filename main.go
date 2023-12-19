package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const luaFilename = "index.lua"

func mainProgram() int {
	// TODO: Is this needed with glfw?
	runtime.LockOSThread()

	L := InitLua(luaFilename)
	defer L.Close()

	windowTitle, err := GetLuaGlobalString(L, "windowTitle")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s must declare a top level windowTitle variable\n", luaFilename)
		return 1
	}

	window := initGraphics(windowTitle)
	defer glfw.Terminate()

	CallLuaFunction(L, "atStart")

	for !window.ShouldClose() {
		ClearScreen()

		CallLuaFunction(L, "atEveryFrame")

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
