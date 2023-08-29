FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod go.sum main.go ./
RUN go mod download
RUN go build -o main .

FROM alpine
COPY --from=builder /build/main /main
CMD ["/main"]
