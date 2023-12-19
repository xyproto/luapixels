package luapixels

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
	lua "github.com/yuin/gopher-lua"
)

var shouldQuit = false

// Run is for running luapixel Lua code, given as a string
func Run(luaCode string) error {
	runtime.LockOSThread()

	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate() // Ensure GLFW is terminated when function exits

	// Create the GLFW window here (or ensure it's created before this function is called)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Initial Title", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create GLFW window: %v", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)

	if err := InitGL(window); err != nil {
		return err
	}

	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("setpal", L.NewFunction(setPalette))
	L.SetGlobal("plot", L.NewFunction(putPixel))
	L.SetGlobal("getpixel", L.NewFunction(getPixel))
	L.SetGlobal("background", L.NewFunction(drawBackground))
	L.SetGlobal("quit", L.NewFunction(quit))

	if err := L.DoString(strings.TrimSpace(luaCode)); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing Lua code: %s\n", err)
		return err
	}

	windowTitle, err := GetLuaGlobalString(L, "window_title")
	if err != nil {
		return errors.New("Lua code must declare a top level windowTitle variable")
	}

	window.SetTitle(windowTitle)

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

// RunFile is for running luapixel Lua code, given a filename
func RunFile(luaFilename string) error {
	data, err := os.ReadFile(luaFilename)
	if err != nil {
		return err
	}
	return Run(string(data))
}
