.PHONY: build

build:
	mewn build -o build/kombo
	mv build/kombo /usr/local/bin/kombo