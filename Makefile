.PHONY: build run

default: build

build:
	@echo "Building the app..."
	@docker build -t go-app .

run: 
	@docker run -it --rm --name gogogo go-app	
