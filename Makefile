pkgname := hyprkeys
version := v1.0.2

build: ${pkgname}

${pkgname}: $(shell find . -name '*.go')
	mkdir -p bin
	go build -o bin/${pkgname} -ldflags "-X 'main.version=${version}'" .

completions:
	mkdir -p completions
	./bin/${pkgname} completion zsh > completions/_${pkgname}
	./bin/${pkgname} completion bash > completions/${pkgname}
	./bin/${pkgname} completion fish > completions/${pkgname}.fish

run:
	go run main.go

tidy:
	go mod tidy

clean:
	rm -rf bin
	rm -rf completions

uninstall:
	rm -f /usr/local/bin/${pkgname}
	rm -f /usr/share/zsh/site-functions/_${pkgname}
	rm -f /usr/share/bash-completion/completions/${pkgname}
	rm -f /usr/share/fish/vendor_completions.d/${pkgname}.fish

install:
	cp bin/${pkgname} /usr/local/bin
	bin/${pkgname} completion zsh > /usr/share/zsh/site-functions/_${pkgname}
	bin/${pkgname} completion bash > /usr/share/bash-completion/completions/${pkgname}
	bin/${pkgname} completion fish > /usr/share/fish/vendor_completions.d/${pkgname}.fish
