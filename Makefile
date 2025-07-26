.PHONY: run
run: build
	go run .

.PHONY: build
build: ignore/roboto/LICENSE ignore/font/ttf-symbola-8.00/LICENSE
	go build .

ignore/roboto/LICENSE: ignore/roboto-unhinted.zip
	unzip ignore/roboto-unhinted.zip -d ignore/roboto
	touch ignore/roboto/LICENSE

ignore/roboto-unhinted.zip: ignore/stamp
	curl -L https://github.com/googlefonts/roboto-2/releases/download/v2.138/roboto-unhinted.zip -o ignore/roboto-unhinted.zip

ignore/font/ttf-symbola-8.00/LICENSE: ignore/symbola.tar.gz ignore/font/stamp
	tar xvzf ignore/symbola.tar.gz --directory=ignore/font
	touch ignore/font/ttf-symbola-8.00/LICENSE

ignore/symbola.tar.gz: ignore/stamp
	curl -L https://github.com/ChiefMikeK/ttf-symbola/archive/refs/tags/v8.00.tar.gz -o ignore/symbola.tar.gz

ignore/font/stamp: ignore/stamp
	mkdir ignore/font
	touch ignore/font/stamp

ignore/stamp:
	mkdir ignore
	touch ignore/stamp
