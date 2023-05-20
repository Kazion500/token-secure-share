build:
	go build -o ./tmp/main main.go

run:
	air

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

