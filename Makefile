build:
	@go build  -o todo-cli
dev: build
	@./todo-cli