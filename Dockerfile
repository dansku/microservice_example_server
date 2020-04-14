# Let's use the last golang base image
FROM golang:latest

# Add maintainer info
LABEL maintainer="Daniel Andrade <daniel@danielandrade.net>"

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port to the outside world based on .env file
EXPOSE ${SERVER_PORT}

# Command to run the executable
CMD ["./main"]