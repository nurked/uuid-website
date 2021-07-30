FROM golang AS build

#ENV GOPRIVATE=github.com/uuid6,github.com/nurked

RUN git clone https://github.com/nurked/uuid-website 
WORKDIR /go/uuid-website
RUN go mod download
RUN go get github.com/vugu/vugu
RUN go get -u github.com/vugu/vugu/cmd/vugugen 
RUN go install github.com/vugu/vugu/cmd/vugugen 
RUN go build -o /go/bin/vugugen github.com/vugu/vugu/cmd/vugugen
RUN go generate
#RUN ls
RUN GOOS=js GOARCH=wasm go build -o output.wasm .
#RUN go build

FROM golang AS prod
RUN mkdir /go/site
WORKDIR /go/site
COPY --from=build /go/uuid-website/output.wasm .
COPY --from=build /go/uuid-website/prodserver.go .
COPY --from=build /go/uuid-website/go.mod .
RUN go mod download github.com/vugu/vugu
RUN go build prodserver.go
RUN go install prodserver.go

CMD ["prodserver"]