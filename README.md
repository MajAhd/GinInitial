# GinInitial
Go Gin Framework Blueprint

## Features
- docker & makefile for easy development
- structured logging with slog


# Scripts 
To see a beautiful, complete list of commands (like running tests, building production binaries, formatting, or linting code), just run:

```bash
make help
```

# Initializing the project

If you'd like to understand how this project was initialized natively or want to replicate the setup from scratch, here are the exact commands and dependencies used to create this blueprint:

```bash
# Initialize the Go module
go mod init gininitial

# Install the primary framework: Gin
go get -u github.com/gin-gonic/gin

# Install environment variable loader
go get github.com/joho/godotenv

# Clean up and finalize the downloaded dependencies
go mod tidy
```

No external logging libraries (such as `logrus` or `zerolog`) were used. We integrated Go's native standard library system **`log/slog`** with a custom logging middleware in `main.go`. This keeps the final binary incredibly small and lightweight while still achieving production-grade JSON logging.


## ORM
BUN ORM https://bun.uptrace.dev/guide/golang-orm.html#quick-start
