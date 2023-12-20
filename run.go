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
		return fmt.Errorf("failed to initialize OpenGL: %v", err)
	}

	if err := InitAudio(); err != nil {
		return fmt.Errorf("failed to initialize audio: %v", err)
	}

	defer context.Free()

	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("setpal", L.NewFunction(setPalette))
	L.SetGlobal("plot", L.NewFunction(putPixel))
	L.SetGlobal("getpixel", L.NewFunction(getPixel))
	L.SetGlobal("background", L.NewFunction(drawBackground))
	L.SetGlobal("quit", L.NewFunction(quit))
	L.SetGlobal("playsound", L.NewFunction(playSound))
	L.SetGlobal("sleep", L.NewFunction(sleep))

	if err := L.DoString(strings.TrimSpace(luaCode)); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing Lua code: %s\n", err)
		return err
	}

	windowTitle, err := GetLuaGlobalString(L, "window_title")
	if err != nil {
		return errors.New("Lua code must declare a top level windowTitle variable")
	}

	window.SetTitle(windowTitle)

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		updateViewportAndProjection(w)
	})

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		L.SetGlobal("last_key", lua.LNumber(key))
		pressed := action == glfw.Press
		released := action == glfw.Release
		if pressed || released {
			switch key {
			case glfw.KeyUp, glfw.KeyW:
				if pressed {
					if err := CallLuaFunction(L, "at_up_pressed"); err != nil {
						CallLuaFunction(L, "at_key_pressed")
					}
				} else {
					if err := CallLuaFunction(L, "at_up_released"); err != nil {
						CallLuaFunction(L, "at_key_released")
					}
				}
			case glfw.KeyDown, glfw.KeyS:
				if pressed {
					if err := CallLuaFunction(L, "at_down_pressed"); err != nil {
						CallLuaFunction(L, "at_key_pressed")
					}
				} else {
					if err := CallLuaFunction(L, "at_down_released"); err != nil {
						CallLuaFunction(L, "at_key_released")
					}
				}
			case glfw.KeyLeft, glfw.KeyA:
				if pressed {
					if err := CallLuaFunction(L, "at_left_pressed"); err != nil {
						CallLuaFunction(L, "at_key_pressed")
					}
				} else {
					if err := CallLuaFunction(L, "at_left_released"); err != nil {
						CallLuaFunction(L, "at_key_released")
					}
				}
			case glfw.KeyRight, glfw.KeyD:
				if pressed {
					if err := CallLuaFunction(L, "at_right_pressed"); err != nil {
						CallLuaFunction(L, "at_key_pressed")
					}
				} else {
					if err := CallLuaFunction(L, "at_right_released"); err != nil {
						CallLuaFunction(L, "at_key_released")
					}
				}
			case glfw.KeySpace:
				if pressed {
					if err := CallLuaFunction(L, "at_space_pressed"); err != nil {
						CallLuaFunction(L, "at_key_pressed")
					}
				} else {
					if err := CallLuaFunction(L, "at_space_released"); err != nil {
						CallLuaFunction(L, "at_key_released")
					}
				}
			case glfw.KeyEnter:
				if pressed {
					if err := CallLuaFunction(L, "at_enter_pressed"); err != nil {
						CallLuaFunction(L, "at_key_pressed")
					}
				} else {
					if err := CallLuaFunction(L, "at_enter_released"); err != nil {
						CallLuaFunction(L, "at_key_released")
					}
				}
			case glfw.KeyEscape:
				if pressed {
					if err := CallLuaFunction(L, "at_esc_pressed"); err != nil {
						CallLuaFunction(L, "at_key_pressed")
					}
				} else {
					if err := CallLuaFunction(L, "at_esc_released"); err != nil {
						CallLuaFunction(L, "at_key_released")
					}
				}
			default:
				if pressed {
					CallLuaFunction(L, "at_key_pressed")
				} else {
					CallLuaFunction(L, "at_key_released")
				}
			}
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
