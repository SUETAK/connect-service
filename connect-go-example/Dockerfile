# Use the official Golang image to create a build artifact.
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o main ./cmd/main.go

# Use the official Alpine image for a lean production container.
FROM alpine:3.11
RUN apk --no-cache add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/main /main

# Expose port 8080 to the outside world
EXPOSE 8000

ENV PORT 8000

# Run the binary program produced by ‘go install’
CMD ["/main"]
