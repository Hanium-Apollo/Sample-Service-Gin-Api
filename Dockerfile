FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./main.go

FROM alpine:latest
COPY --from=builder /app/main /main
COPY --from=builder /app/.env /.env
CMD ["/main"]
