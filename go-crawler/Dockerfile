FROM golang:1.13-alpine as builder
WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o appbin

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/appbin .
COPY --from=builder /app/log.txt .
CMD ["/app/appbin"]