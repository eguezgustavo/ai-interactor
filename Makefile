format:
	@echo "Formatting files..."
	@gofmt -w $$(find . -type f -name '*.go')
	@echo "OK"
	@echo

test:
	@echo "Formatting files..."
	@go test ./...
	@echo

build/edge:
	@echo "Building for edge devices (MIPS, linux, softfloat)..."
	@GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o ai-interactor
	@echo
