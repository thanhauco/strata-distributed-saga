BIN_DIR := bin

build:
	go build -v ./...

test:
	go test -race -v -cover ./...

lint:
	golangci-lint run

tf-init:
	cd terraform && terraform init

tf-plan:
	cd terraform && terraform plan

clean:
	rm -rf $(BIN_DIR) *.out

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
