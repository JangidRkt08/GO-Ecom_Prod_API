# --------------------------------------------------------
# Build stage
# --------------------------------------------------------
FROM golang:1.22-alpine AS build

WORKDIR /app

# Install dependencies required for goose + database/sql
RUN apk add --no-cache gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go server
RUN go build -o server ./cmd

# --------------------------------------------------------
# Run stage
# --------------------------------------------------------
FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy compiled binary
COPY --from=build /app/server /app/server

# Copy migrations for Goose (IMPORTANT!)
COPY internal/adapters/postgresql/migrations /app/migrations

EXPOSE 8080

CMD ["./server"]
