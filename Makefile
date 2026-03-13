SCRIPTS_DIR ?= $(HOME)/Development/github.com/rios0rios0/pipelines
-include $(SCRIPTS_DIR)/makefiles/common.mk
-include $(SCRIPTS_DIR)/makefiles/golang.mk

.PHONY: build build-musl debug install run

build:
	rm -rf bin
	go build -o bin/ronin-to-koinly ./cmd/ronin-to-koinly
	strip -s bin/ronin-to-koinly

debug:
	rm -rf bin
	go build -gcflags "-N -l" -o bin/ronin-to-koinly ./cmd/ronin-to-koinly

build-musl:
	CGO_ENABLED=1 CC=musl-gcc go build \
		--ldflags 'linkmode external -extldflags="-static"' -o bin/ronin-to-koinly ./cmd/ronin-to-koinly
	strip -s bin/ronin-to-koinly

run:
	go run ./cmd/ronin-to-koinly

install:
	make build
	mkdir -p ~/.local/bin
	cp -v bin/ronin-to-koinly ~/.local/bin/ronin-to-koinly
