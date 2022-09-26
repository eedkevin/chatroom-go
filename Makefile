.PHONY: install
install:
	go install github.com/cosmtrek/air@latest

.PHONY: run
run: 
	go run main.go

.PHONY: watch
watch:
	air -c .air.toml