.PHONY: install
install:
	go install github.com/cosmtrek/air@latest
	go install golang.org/x/lint/golint@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	brew install pre-commit
	pre-commit install

.PHONY: run
run: 
	go run main.go

.PHONY: watch
watch:
	air -c .air.toml