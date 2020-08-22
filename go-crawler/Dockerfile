FROM golang:1.14.3-alpine
WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o appbin cmd/crawl-hunterio/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/appbin .
COPY --from=0 /app/static ./static
CMD ["/app/appbin"]