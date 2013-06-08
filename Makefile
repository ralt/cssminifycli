all:
	go build -o dist/cssminify

install:
	cp dist/cssminify /usr/bin
