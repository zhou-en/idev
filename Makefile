hello:
	echo "Hello"

build:
	go build -o bin/idev main.go

run:
	# make run ARGS="--parJson '{\\\"a\\\":\\\"aa\\\"}'"
	go build -o bin/idev main.go && ./bin/idev $(ARGS)

release:
	cp bin/idev $(GOPATH)/bin

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
