all: build

build:
	go build -o bin/caligula

run: build
	./bin/caligula
