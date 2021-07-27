FROM golang AS build

ENV GOOS=js
ENV GOARCH=wasm
ENV GOPRIVATE=github.com/uuid6,github.com/nurked

RUN git clone https://github.com/nurked/uuid-website 
WORKDIR uuid-website 
RUN go mod download
RUN go generate 
RUN go build