package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

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

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		L.SetGlobal("last_key", lua.LNumber(key))
		if action == glfw.Press {
			CallLuaFunction(L, "at_keypress")
		} else if action == glfw.Release {
			CallLuaFunction(L, "at_keyrelease")
		}
	})

	CallLuaFunction(L, "at_start")

	tick := time.NewTicker(time.Millisecond * 16) // 60 ticks per second
	defer tick.Stop()

	for !window.ShouldClose() && !shouldQuit {
		select {
		case <-tick.C:
			CallLuaFunction(L, "at_every_tick")
		default:
			// Non-blocking default case
		}

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
