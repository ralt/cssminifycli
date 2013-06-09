package main

import (
	"cssminify"
)

func main() {
	files := cssminify.Files()
	f := make(chan string)
	b := make(chan []cssminify.Block)
	go cssminify.Blocks(f, b)
	go cssminify.Minify(b)
	for _, file := range files {
		f <- file
	}
}
