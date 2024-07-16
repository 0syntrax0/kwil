# syntax=docker/dockerfile:1

# Start from golang base image
FROM golang:1.22-alpine

# Add Maintainer info
LABEL maintainer="Carlos Saavedra <c.saavedra85@hotmail.com>"

# Set the current working directory inside the container 
WORKDIR /api

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the executable
RUN go build -o ./build/dist .

# Run the executable
CMD ./build/dist
