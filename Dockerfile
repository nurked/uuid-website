FROM golang AS build

ENV GOOS=js
ENV GOARCH=wasm

RUN git clone https://github.com/nurked/uuid-website && \
    cd uuid-website && \
    go mod download github.com/uuid6/uuid6go-proto && \
    go mod download github.com/vugu/vugu && \
    go mod download github.com/vugu/vjson && \
    go mod download github.com/vugu/vgrouter && \
    go generate && \
    go build