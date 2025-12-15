setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	make install

install:
	go mod download

brain-games:
	@go run cmd/brain_games/main.go

brain-even:
	@go run cmd/brain_even/main.go

brain-calc:
	@go run cmd/brain_calc/main.go

brain-gcd:
	@go run cmd/brain_gcd/main.go

brain-progression:
	@go run cmd/brain_progression/main.go

brain-prime:
	@go run cmd/brain_prime/main.go

check:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

test:
	go test ./... -v
