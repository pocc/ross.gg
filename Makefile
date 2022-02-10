.PHONY: build

build:
	hugo --theme=even -D

deploy: build
	mv public/* /var/www/html
