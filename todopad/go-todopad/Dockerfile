# Stage 1: Build the Go binary with CGO enabled
FROM golang:1.23-bookworm AS build

# Set working directory inside container
WORKDIR /app

# Install required dependencies for CGO and SQLite
RUN apt-get update && apt-get install -y gcc g++ make sqlite3 libsqlite3-dev

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Enable CGO and build the Go binary
ENV CGO_ENABLED=1
RUN go build -o /app/main .

# Stage 2: Create a minimal runtime image
FROM debian:bookworm-slim

# Set working directory
WORKDIR /app

# Install SQLite runtime dependencies
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Copy the compiled Go binary
COPY --from=build /app/main .

# Copy the database file (if it exists)
COPY --chown=1000:1000 todos.db /app/todos.db

# Copy templates and static files
COPY templates/ /app/templates/

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./main"]
