FROM golang:1.25.4 AS builder
WORKDIR /app
COPY main.go .
COPY go.mod .

RUN <<EOF
go mod download
go mod tidy
# CGO_ENABLED=0 creates a statically linked binary without dynamic dependencies
# This is required for running the binary in a scratch image
CGO_ENABLED=0 go build -o tiny-service
EOF

FROM scratch
WORKDIR /app
COPY --from=builder /app/tiny-service .

COPY public ./public 
CMD ["./tiny-service"]

