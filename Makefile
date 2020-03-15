.PHONY: build

build:
	go build -o build/kombo
	mv build/kombo /usr/local/bin/kombo