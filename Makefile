ARCH ?= amd64

build-darwin: librobot.go
	env GOOS=darwin GOARCH=$(ARCH) go build -buildmode=c-shared -o librobot.dylib librobot.go

build-linux: librobot.go
	env GOOS=linux GOARCH=$(ARCH) go build -buildmode=c-shared -o librobot.so librobot.go

clean:
	rm -f librobot.dylib librobot.so
