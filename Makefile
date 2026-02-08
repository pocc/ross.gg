.PHONY: build dev clean

build:
	hugo --minify

dev:
	hugo server --buildDrafts

clean:
	rm -rf public/ resources/
