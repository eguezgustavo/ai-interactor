format:
	@echo "Formatting files..."
	@gofmt -w $$(find . -type f -name '*.go')
	@echo "OK"
	@echo

test:
	@echo "Formatting files..."
	@go test ./...
	@echo

build:
	@echo "Building..."
	@go build -o ai-interactor
	@echo "OK"
	@echo

build/edge:
	@echo "Building for edge devices (MIPS, linux, softfloat)..."
	@GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o ai-interactor
	@echo "OK"
	@echo

deploy:
	@echo "Uploading ai-interactor to $$DEVICE_URL..."
	@scp ai-interactor $$USER@$$DEVICE_URL:/usr/bin
	@echo "OK"
	@echo

install-deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@echo "OK"
	@echo