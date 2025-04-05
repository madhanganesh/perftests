# Build stage
FROM rust:bookworm AS builder

WORKDIR /app
COPY . .
ENV DATABASE_URL=sqlite:///app/todos.db

RUN cargo build --release

# Final run stage
FROM debian:bookworm-slim AS runner

# Install necessary system dependencies
RUN apt-get update && apt-get install -y libssl3 sqlite3 ca-certificates && apt-get clean

WORKDIR /app

# Copy compiled Rust binary
COPY --from=builder /app/target/release/ru-todopad /app/todopad

# Copy database migrations and static files
COPY migrations /app/migrations

# Copy the SQLite database file into the image
COPY todos.db /app/todos.db

# Set the DATABASE_URL environment variable
ENV DATABASE_URL=sqlite:///app/todos.db

EXPOSE 8080

# Run the application
CMD ["/app/todopad"]
