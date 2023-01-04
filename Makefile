BINARY_NAME=hyprkeys

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} main.go

run:
	go run main.go

tidy:
	go mod tidy

clean:
	rm -rf bin

install: build inst

uninstall:
	rm -f /usr/local/bin/${BINARY_NAME}

inst:
	cp bin/${BINARY_NAME} /usr/local/bin


