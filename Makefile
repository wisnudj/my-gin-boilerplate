NAME=my-gin-boilerplate
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME)

.PHONY: clean
clean:
	@rm -f $(NAME)