package main

import (
	"log"

	"github.com/xyproto/luapixels"
)

func main() {
	if err := luapixels.RunFile("events.lua"); err != nil {
		log.Fatalln(err)
	}
}
