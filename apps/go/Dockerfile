FROM golang:1.21 as builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum /app/
RUN go mod download

# Copy the rest of the application
COPY . /app

# Build the Go application binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy compiled binary and wait-for-it script
COPY --from=builder /app/main .
COPY wait-for-it.sh /wait-for-it.sh

# Grant execute permissions to wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 8080
ENTRYPOINT ["/wait-for-it.sh", "postgres:5432", "--", "mongodb:27017", "--", "./main"]
