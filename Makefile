build:
	go build -v -ldflags "-X main.commit=$(shell git log --pretty=format:"%h" -1)"

install:
	sudo cp tool /usr/local/bin

all: build install
