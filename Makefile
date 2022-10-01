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

.PHONY: compose-up
compose-up:
	docker compose up --build

.PHONY: compose-down
compose-down:
	docker compose down

.PHONY: test
test:
	go test -v ./...

.PHONY: client
client:
	go run pkg/client/client.go -user ${user} -room ${room}