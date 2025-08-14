# Stage 1: Build the go applcation
FROM golang:alpine AS builder

WORKDIR /app

# Copy the Go module manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download && go mod verify

# Copy the source code
COPY . .

# Build Go App
RUN CGO_ENABLED=0 GOOS=linux GOARCHH=amd64 go build -a -installsuffix cgo -o user-service ./cmd

# Stage 2: Create a minimal image
FROM alpine:latest

# Set the working dir
WORKDIR /app

# Add some necessary packages
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

# Copy the binary from the builder stage
COPY --from=builder /app/user-service .

# Copy the environtment file
COPY --from=builder /app/.env .

# Copy entrypoint sh
COPY --from=builder /app/script/entrypoint.sh .

# Expose the port the app runs on
EXPOSE 8080

# Wait-for-it requires bash, which alpine does'nt ship with by default.
# Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for ./entrypoint.sh

# Command to run the binary
ENTRYPOINT [ "sh", "./entrypoint.sh" ]