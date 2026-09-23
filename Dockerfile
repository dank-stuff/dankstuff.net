FROM golang:1.26-alpine AS build

WORKDIR /app
COPY . .

RUN apk add --no-cache wget make sed git npm nodejs
RUN go install github.com/a-h/templ/cmd/templ@v0.3.1020
RUN make build-http-server

FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app/dankstuff.net .

EXPOSE 8080

CMD ["./dankstuff.net"]
