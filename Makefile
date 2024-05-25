include .env
export

build:
	go build -o bot

run: build
	./bot