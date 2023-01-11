BINARY_NAME=hyprkeys

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} main.go

run:
	go run main.go

tidy:
	go mod tidy

clean:
	rm -rf bin

uninstall:
	rm -f /usr/local/bin/${BINARY_NAME}

install:
	cp bin/${BINARY_NAME} /usr/local/bin


