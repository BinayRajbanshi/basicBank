# 🏦 GoBank – Backend API with Go, PostgreSQL & Gin

A production-grade backend service written in Go, demonstrating best practices in API development, database integration, testing, and concurrency. This project is designed for performance, scalability, and reliability.

## 🚀 Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** Gin
- **Database:** PostgreSQL
- **Token Auth:** PASETO & JWT
- **Containerization:** Docker
- **CI/CD:** GitHub Actions

## 📦 Features

### 🗄️ Robust Database Layer

- Clean database schema design using PostgreSQL
- Auto-generated database access code using sqlc
- Safe and consistent queries with transaction support
- Proper use of isolation levels to prevent concurrency issues

### 🧪 Testing and CI

- Comprehensive unit and integration test coverage
- Database mocking for isolated and reliable test cases
- Automated test runs using GitHub Actions on every push

### 🔐 Secure REST APIs

- RESTful API design using the Gin framework
- User authentication and authorization
- Secured endpoints using JWT and PASETO tokens
- Structured error handling and configuration management

### 🐳 Local Development with Docker

- Full Docker support for local development and testing
- Isolated environment for the database and app server
- Easy spin-up with docker-compose

### ⚙️ Concurrency with Go

Go's built-in support for concurrency using goroutines and channels enables this backend to handle thousands of concurrent requests with minimal resource consumption. Unlike traditional multi-threaded models found in languages like Java, Python, or Node.js (JavaScript), Go achieves concurrency with lightweight green threads managed by the Go runtime, not the OS.

This gives Go a distinct performance and scalability edge over:

- **Java + Spring Boot** - more memory-hungry and complex thread management
- **Python + Django/Flask** - limited by the Global Interpreter Lock (GIL)
- **Node.js + Express** - single-threaded, relies heavily on event loop callbacks
- **Ruby on Rails** - less performant in high-concurrency environments

Go's model simplifies concurrent programming and avoids many pitfalls like thread locks, race conditions, and deadlocks - making it ideal for high-performance backend services and microservices.

## 🔧 Setup & Makefile Commands

All setup and management tasks are streamlined using a Makefile. Refer to the targets below for quick access:

| Target | Description |
|--------|-------------|
| `postgres` | Start a local PostgreSQL 17 container |
| `createdb` | Create the basic_bank database |
| `dropdb` | Drop the basic_bank database |
| `migrateup` | Apply all up migrations |
| `migratedown` | Roll back the last migration |
| `sqlc` | Generate Go code from SQL queries |
| `test` | Run all unit and integration tests |
| `server` | Start the application server |
| `mock` | Generate mocks for database interfaces |

To use any of the above, simply run:

```bash
make <target>
```

**Example:**
```bash
make postgres
```

> ℹ️ The Makefile also uses `.PHONY` to ensure commands are always executed, not treated as file targets.

## 🧪 How to Run Tests

```bash
make test
```

## 🐳 How to Run with Docker

```bash
make postgres
make createdb
make migrateup
make server
```