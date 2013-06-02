package main

import (
	"cssminify"
	"sync"
)

func main() {
	files := cssminify.Files()
	cb := make(chan cssminify.Block)

	var wg sync.WaitGroup
	wg.Add(1)
	go getBlocks(cb, files, wg)
	wg.Wait()

	wg.Add(1)
	go cssminify.Minify(cb, wg)
	wg.Wait()
}

func getBlocks(cb chan cssminify.Block, files []string, wg sync.WaitGroup) {
	for _, file := range files {
		cssminify.Blocks(cb, file, wg)
	}
	defer wg.Done()
}
