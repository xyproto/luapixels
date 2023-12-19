package main

import (
	"log"

	_ "embed"

	"github.com/xyproto/luapixels"
)

//go:embed index.lua
var luaCode string

func main() {
	if err := luapixels.Run(luaCode); err != nil {
		log.Fatalln(err)
	}
}
