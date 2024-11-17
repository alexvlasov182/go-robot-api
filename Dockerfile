# Use the  official Golang image as the base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory into the container
COPY . ./

# Build the Go application
RUN go build -o main .

# Run the binary
CMD ["./main"]
