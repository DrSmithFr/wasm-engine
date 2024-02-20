FROM golang:1.22-alpine as webassembly-builder

RUN mkdir /app
ADD . /app
WORKDIR /app

ENV PORT 8080
EXPOSE 8080

ENV GOOS js
ENV GOARCH wasm


CMD ["go", "run", "main.go"]

FROM
