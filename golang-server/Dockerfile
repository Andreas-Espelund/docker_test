# GO Base image
FROM golang:1.16

# Set directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download and verify go modules
RUN go mod download && go mod verify

# Copy source code into container
COPY . .

# Build the go program
RUN go build -o main .

# set entry point of container

CMD ["./main"]

EXPOSE 8080
