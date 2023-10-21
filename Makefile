build:
	go build -v

install:
	sudo cp tool /usr/local/bin

all: build install
