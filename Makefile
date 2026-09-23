.PHONY: all build-http-server

BINARY_NAME=dankstuff.net

TEMPL_CMD=templ
ifdef CI
	TEMPL_CMD := go run github.com/a-h/templ/cmd/templ@v0.3.1020
endif

all: build-http-server

build-http-server: generate
	@go build -ldflags="-w -s" -o ${BINARY_NAME} ./cmd/http/main.go

init: go-init tailwindcss-init

generate:
	@${TEMPL_CMD} generate -path .

go-init:
	@go mod tidy

tailwindcss-init:
	@mkdir -p assets/css &&\
	npm i &&\
	npx @tailwindcss/cli -i assets/css/style.css -o assets/css/tailwind.css -m

tailwindcss-build:
	@npx @tailwindcss/cli -i assets/css/style.css -o assets/css/tailwind.css

tailwindcss-server:
	@npx @tailwindcss/cli -i assets/css/style.css -o assets/css/tailwind.css --watch

local-http-server: build-http-server
	@./${BINARY_NAME}

clean:
	@rm -f ./${BINARY_NAME}*
	@go clean

