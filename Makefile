.PHONY: build clean

build:
	go build -o bin/ctm ./cmd/main.go

clean:
	rm -rf bin/*

install: build
	cp bin/ctm /usr/local/bin/  # or another location in your PATH

test:
	go test ./...
