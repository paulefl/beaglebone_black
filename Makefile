CROSS   = arm-linux-gnueabihf-
CC      = $(CROSS)gcc
TARGET  = armv7-unknown-linux-musleabihf

all: c-lib rust-lib go-api

c-lib:
	$(MAKE) -C c-lib CC=$(CC)
	cp c-lib/libhardware.so go-api/libs/ 2>/dev/null || true

rust-lib:
	cd rust-lib && cross build --release --target $(TARGET)
	cp rust-lib/target/$(TARGET)/release/libhardware_rs.so go-api/libs/ 2>/dev/null || true
	cbindgen --config rust-lib/cbindgen.toml --output go-api/libs/include/hardware_rs.h 2>/dev/null || true

go-api: c-lib rust-lib
	cd go-api && GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=$(CC) \
	  go build -ldflags="-s -w" -o ../bin/embedded ./cmd/

cli:
	cd tools/cli && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	  go build -o ../../bin/bbcli-linux-amd64 .

test:
	./scripts/test.sh

test-ci:
	./scripts/test.sh -ci

test-cover:
	./scripts/test.sh -cover -html

lint:
	cd go-api && go vet ./pkg/hal/ ./pkg/hal/mock/ ./pkg/hal/config/
	test -z "$$(gofmt -l go-api/ tools/)" || (echo "❌ Formatierung prüfen: gofmt -w ." && exit 1)
	@echo "✅ Lint OK"

deploy:
	scp bin/embedded go-api/libs/libhardware.so go-api/libs/libhardware_rs.so \
	  debian@192.168.7.2:/app/
	ssh debian@192.168.7.2 "systemctl restart embedded-sw"

clean:
	$(MAKE) -C c-lib clean
	cd rust-lib && cargo clean
	rm -f bin/embedded bin/bbcli-*

.PHONY: all c-lib rust-lib go-api cli test test-ci test-cover lint deploy clean
