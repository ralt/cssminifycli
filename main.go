package main

import (
	"cssminify"
)

func main() {
	files := cssminify.Files()
	for _, file := range files {
		cssminify.Minify(cssminify.Blocks(file), file)
	}
}
