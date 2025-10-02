BINARY=equinix

build:
	go build $(LDFLAGS) -o bin/$(BINARY) main.go