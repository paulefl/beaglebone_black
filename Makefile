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
	cd tools/cli && go mod tidy && go vet ./...
	cd tools/tui && go mod tidy && go vet ./...
	test -z "$$(gofmt -l go-api/ tools/)" || (echo "❌ Formatierung prüfen: gofmt -w ." && exit 1)
	@echo "✅ Lint OK"

test-python:
	pytest tests/api/ -v --timeout=10

test-report:
	./scripts/report.sh

test-report-open:
	./scripts/report.sh --open

deploy:
	scp bin/embedded go-api/libs/libhardware.so go-api/libs/libhardware_rs.so \
	  debian@192.168.7.2:/app/
	ssh debian@192.168.7.2 "systemctl restart embedded-sw"

req-tracing:
	strictdoc --debug export . --formats html     --output-dir output/strictdoc
	strictdoc --debug export . --formats html2pdf --output-dir output/strictdoc
	strictdoc --debug export . --formats excel    --output-dir output/strictdoc
	strictdoc --debug export . --formats reqif-sdoc --output-dir output/strictdoc
	python3 scripts/req_tracing_summary.py

clean:
	$(MAKE) -C c-lib clean
	cd rust-lib && cargo clean
	rm -f bin/embedded bin/bbcli-*

.PHONY: all c-lib rust-lib go-api cli test test-ci test-cover lint test-python test-report test-report-open deploy clean req-tracing
