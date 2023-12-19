package luapixels

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
	lua "github.com/yuin/gopher-lua"
)

var shouldQuit = false

// Run is for running luapixel Lua code, given as a filename
func Run(luaFilename string) error {
	data, err := os.ReadFile(luaFilename)
	if err != nil {
		return err
	}
	return RunCode(string(data))
}

// RunCode is for running luapixel Lua code, given as a string
func RunCode(luaCode string) error {
	runtime.LockOSThread()

	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(luaCode); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing Lua code: %s\n", err)
		return err
	}

	windowTitle, err := GetLuaGlobalString(L, "window_title")
	if err != nil {
		return errors.New("Lua code must declare a top level windowTitle variable")
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
	return nil
}
