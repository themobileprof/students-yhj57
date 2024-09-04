# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Define build-time variable
ARG DB=test.db

# Set the environment variable
ENV DB=${DB}

# Build the Go application
RUN go build -o student-management-api main.go

# Expose the port the application runs on
EXPOSE 8080
# Command to run the application
CMD ["./student-management-api"]
