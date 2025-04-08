# Concurrent In-Memory Cache Server

## Overview

This project implements a lightweight, concurrent in-memory cache server built in Go. The server exposes a RESTful API for basic cache operations (set, get, delete) and tracks cache metrics (hits, misses, item count). It also automatically evicts expired cache entries using a background TTL (Time-To-Live) manager.

This project was developed to demonstrate core computer science concepts such as concurrency, synchronization, REST API design, and basic logging—using Go's native libraries. It serves as a product-grade component that can be integrated into larger systems or used as a standalone service.

## Features Implemented So Far

- **In-Memory Cache Module:**  
  - Stores key–value pairs.
  - Supports an optional TTL per entry.
  - Uses a `sync.RWMutex` for thread-safe access.
  - Includes functions to set, get, delete items, and return cache metrics.
  
- **TTL Manager / Background Cleaner:**  
  - A goroutine running with a `time.Ticker` that periodically removes expired cache entries.
  - Logs eviction events for observability.
  
- **HTTP Server with RESTful API Endpoints:**  
  - Built with Go's `net/http` package.
  - **Endpoints:**
    - `POST /cache/set` – Stores a key–value pair (with an optional TTL).
    - `GET /cache/get?key=...` – Retrieves a value by its key.
    - `DELETE /cache/delete?key=...` – Deletes a specified cache entry.
    - `GET /cache/metrics` – Returns cache metrics (hits, misses, current item count).

- **Project Structure & Environment Setup:**  
  - Proper project layout using Go modules.
  - Version-controlled with Git.
  - Organized code in modular folders: `cmd/`, `internal/api/`, `internal/cache/`.

## Project Structure

```bash
cache-server/
├── cmd/
│   └── main.go                # HTTP server entry point
├── internal/
│   ├── api/
│   │   └── handlers.go        # REST API handlers for cache operations
│   └── cache/
│       ├── cache.go           # In-memory cache implementation & metrics
│       └── cleanup.go         # TTL manager/background cleaner
├── tests/                     # (Optional) Unit tests will reside here
├── go.mod                     # Go module file
```

## Getting Started

### Prerequisites

- **Go 1.18+** installed on your machine.
- A text editor or IDE (e.g., Visual Studio Code with the Go extension).
- Git (optional, for version control).

### Environment Setup Summary

1. **Installed Go** and verified with `go version`.
2. **Set up your IDE,** and installed necessary Go extensions.
3. **Created the project directory,** initialized a Git repository, and set up the Go module:

   ```bash
   mkdir cache-server && cd cache-server
   git init
   go mod init github.com/yourusername/cache-server
   ```

4. **Organized the project structure** into folders for commands, internal modules, and tests.

## How to Run the Application

1. **Build and Run:** From the project root, run:

   ```bash
   go run cmd/main.go -port=8080
   ```

   The HTTP server will start and listen on the specified port (default is `8080`).

2. **Testing the Endpoints using Curl:**
   - **Set a Cache Entry:**

     ```bash
     curl -X POST -d '{"key": "example", "value": "hello", "ttl": 10}' http://localhost:8080/cache/set
     ```

   - **Get a Cache Entry:**

     ```bash
     curl "http://localhost:8080/cache/get?key=example"
     ```

   - **Delete a Cache Entry:**

     ```bash
     curl -X DELETE "http://localhost:8080/cache/delete?key=example"
     ```

   - **Get Cache Metrics:**

     ```bash
     curl "http://localhost:8080/cache/metrics"
     ```

## What Has Been Achieved

- **Core Functionality:** An in-memory cache with TTL support that safely handles concurrent access.
- **Automatic Cleanup:** A background TTL manager is now removing expired cache entries automatically.
- **RESTful API:** A production-style HTTP server with endpoints for adding, retrieving, deleting data, and checking metrics.
- **Logging and Basic Metrics:** Metrics are tracked and logged, demonstrating observability in the system.

## To-Do / Next Steps

1. **Unit Testing:**
   - Write unit tests for cache operations (Set, Get, Delete, CleanupExpired).
   - Write tests for the API handlers to ensure proper JSON responses and error handling.

2. **Advanced Logging & Error Handling:**
   - Enhance the logging mechanism for deeper insights.
   - Implement centralized error handling and richer error responses if needed.

3. **Containerization:**
   - Create a `Dockerfile` to package the application.
   - Set up CI/CD pipelines for automated testing and deployment.

4. **Advanced Metrics and Monitoring:**
   - Consider integrating with tools like Prometheus for detailed monitoring.
   - Expose advanced metrics or integrate with centralized logging systems.

5. **Security Enhancements:**
   - Implement input sanitization and rate limiting on API endpoints.
   - Plan for HTTPS deployment in a production environment.

6. **Documentation and Code Comments:**
   - Continue improving inline code documentation.
   - Update this README with any new instructions as the project evolves.
