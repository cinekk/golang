.PHONY: build run test

default: build

current_dir = $(shell pwd)

build:
	@echo "Building the app..."
	@docker build -t go-app .

run: 
	@docker run -it --rm --name gogogo go-app	

test:
	@docker run -it --rm -v $(current_dir):/usr/src/myapp -w /usr/src/myapp golang go test
