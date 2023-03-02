# Use an official Go runtime as a parent image
FROM golang:1.17.2-alpine3.14

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Define the command to run the executable
CMD ["./main"]
