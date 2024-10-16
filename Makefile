build:
	go build -v -ldflags "-X 'main.Commit=$(shell git log --pretty=format:%H -1)'"

install:
	sudo cp tool /usr/local/bin

all: build install
