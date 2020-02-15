FROM golang:1.13-alpine as builder
WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o goapp

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/goapp .
CMD ["/app/goapp"]