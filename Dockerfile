FROM golang:alpine

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARH=amd64 GOMODULE=on go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 8888

CMD ["/dist/main"]