FROM golang:1.24-alpine

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod ./

# Copy the source code and other necessary files
COPY . .

# Build the application
RUN go build -o main .

# Make sure the static and templates directories are included
RUN mkdir -p /app/static /app/templates

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]