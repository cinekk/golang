.PHONY: build run test

default: build

build:
	@echo "Building the app..."
	@docker build -t go-app .

run: 
	@docker run -it --rm --name gogogo go-app	

test:
	@docker run -it --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang sh
