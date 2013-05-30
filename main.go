package main

import (
	"cssminify"
	"fmt"
)

func main() {
	files := cssminify.Files()
	for _, file := range files {
		fmt.Printf("%s", cssminify.Minify(cssminify.Blocks(file)))
	}
}
