package main

import (
	"log"

	"github.com/xyproto/luapixels"
)

func main() {
	if err := luapixels.RunFile("game.lua"); err != nil {
		log.Fatalln(err)
	}
}
