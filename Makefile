.PHONY: build dev clean deploy preview

build:
	hugo --minify

dev:
	hugo server --buildDrafts

clean:
	rm -rf public/ resources/

deploy:
	npx wrangler deploy

preview:
	hugo --minify && npx wrangler dev
