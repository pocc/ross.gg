.PHONY: build

build:
	hugo --theme=even -D

deploy: build
	sudo cp -r public/* /var/www/html
