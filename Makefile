BINARY_NAME=driveshare-backend
.PHONY :  build
dep:
	go mod download

build:
	GOARCH=amd64 GOOS=darwin go build -o ./build/${BINARY_NAME}-darwin ./cmd/app/main.go
	GOARCH=amd64 GOOS=linux go build -o ./build/${BINARY_NAME}-linux ./cmd/app/main.go
	GOARCH=amd64 GOOS=windows go build -o ./build/${BINARY_NAME}-windows ./cmd/app/main.go

run: build
	./build/${BINARY_NAME}-darwin

clean:
	go clean
	rm ./build/${BINARY_NAME}-darwin
	rm ./build/${BINARY_NAME}-linux
	rm ./build/${BINARY_NAME}-windows

test:
	go test ./...