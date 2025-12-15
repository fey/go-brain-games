install:
	go mod download

brain-games:
	@go run cmd/brain_games/main.go

brain-even:
	@go run cmd/brain_even/main.go

check:
	go vet ./...

fmt:
	go fmt ./...
