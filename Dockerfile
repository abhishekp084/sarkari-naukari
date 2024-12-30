# Stage 1: Build stage
FROM golang:1.22 as base

WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

RUN go mod download

# Copy application source code and templates
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Final stage - Distroless image
FROM gcr.io/distroless/base

# Copy the compiled binary and required files from the build stage
COPY --from=base /app/main /
COPY --from=base /app/template.html /

# Expose the application port
EXPOSE 8000

# Command to run the application
CMD ["/main"]
