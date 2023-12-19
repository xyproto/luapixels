package main

import (
    "runtime"

    "github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
    runtime.LockOSThread()

    window := initGraphics("Pixels!")
    defer glfw.Terminate()

    L := InitLua("main.lua")
    defer L.Close()

    CallLuaFunction(L, "atStart")

    for !window.ShouldClose() {
        ClearScreen()

		CallLuaFunction(L, "everyFrame")

        UpdateScreen(window)
        glfw.PollEvents()
    }

    CallLuaFunction(L, "atEnd")
}
