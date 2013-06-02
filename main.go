package main

import (
	"cssminify"
)

func main() {
	files := cssminify.Files()
	cb := make(chan cssminify.Block)

	go getBlocks(cb, files)
	go cssminify.Minify(cb)
}

func getBlocks(cb chan cssminify.Block, files []string) {
	for _, file := range files {
		cssminify.Blocks(cb, file)
	}
}
