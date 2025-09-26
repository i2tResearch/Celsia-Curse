FROM golang:1.25.1-alpine

# Set the working directory
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Initialize the module and download dependencies
# This replaces the deprecated `go get`
RUN go mod init main && go mod tidy

# Build the Go app
RUN go build -o server .

# Run the executable
CMD ["/app/server"]
