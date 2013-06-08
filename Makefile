all:
	go get github.com/Ralt/cssminify
	go build -o dist/cssminify

install:
	cp dist/cssminify /usr/bin
