# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the templ templates first, then the binary
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o blog ./main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy binary and runtime assets
COPY --from=builder /app/blog .
COPY --from=builder /app/content ./content
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./blog"]
