hello:
	echo "Hello"

build:
	go build -o bin/idev main.go

run:
	go build -o bin/idev main.go && ./bin/idev $(ARGS)

release:
	cp bin/idev /usr/local/bin

clean:
	go clean

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod tidy

lint:
	golangci-lint run --disable=typecheck,gosimple,staticcheck,structcheck,unused ./... cmd/
