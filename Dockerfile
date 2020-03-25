# Start from the latest golang base image
ARG GO_VERSION=1.13.1
FROM golang:${GO_VERSION} AS builder
CMD ["bash"]
# Add Maintainer Info
LABEL maintainer="Orkhan Huseynli"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./ ./


WORKDIR /app/src

# Build the Go app
RUN go build -o main .

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["bash", "-c", "exec ./main"]
