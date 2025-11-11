🧩 90-DAY GO BACKEND MASTERY ROADMAP

(For Backend Jobs + Open-Source Contribution)

🗓️ PHASE 1: Go Language Foundations (Week 1–3)

🎯 Goal: Build a rock-solid command over Go syntax, types, and idioms.

📅 Week 1 — Go Basics & Syntax

✅ Topics:

Installing Go & setting up VSCode

Go workspace, modules (go mod init, go run)

Variables, constants, types, and type inference

Loops, conditionals, functions

Arrays, slices, maps

Pointers (value vs reference types)

💡 Mini Project:
➡️ Build a simple CLI Calculator using functions (add, sub, mul, div)

📅 Week 2 — Structs, Methods & Interfaces

✅ Topics:

Structs and methods

Receivers (value vs pointer)

Interfaces and interface composition

Embedding (struct + interface)

Custom types and type aliases

Error handling (error, fmt.Errorf, custom errors)

💡 Mini Project:
➡️ “Library Management” system using structs (Book, User, Borrower).
➡️ Use interfaces for different roles (Admin, User).

📅 Week 3 — Packages, Modules, and Error Handling

✅ Topics:

Package structure (cmd/, pkg/, internal/)

Defer, Panic, Recover

Go Modules and dependency management

Basic file handling (read/write)

Time and formatting utilities

💡 Mini Project:
➡️ Build a “Task Manager” CLI app that saves tasks to a text file.

🗓️ PHASE 2: Go Concurrency & Backend Building (Week 4–7)

🎯 Goal: Master goroutines, channels, context, and start building REST APIs.

📅 Week 4 — Concurrency Deep Dive

✅ Topics:

Goroutines (spawn, wait)

Channels (unbuffered & buffered)

select statement

WaitGroups & Mutex

Context package (cancel, timeout)

Worker pools pattern

💡 Mini Project:
➡️ Build a “Concurrent Downloader” — multiple file downloads in parallel using goroutines.

📅 Week 5 — net/http & REST API Fundamentals

✅ Topics:

net/http package

Handlers, routing, response writing

JSON encoding/decoding

URL params and query params

Error handling and status codes

Middleware (logging, recovery)

💡 Mini Project:
➡️ Build a “Notes REST API” using native net/http (CRUD for Notes).

📅 Week 6 — Frameworks & Middleware

✅ Topics:

Frameworks: Gin / Fiber / Chi

Router setup, groups, middleware

Validation (binding JSON)

Authentication (JWT)

Environment variables (os.Getenv, godotenv)

Configuration using Viper

💡 Mini Project:
➡️ “User Authentication API” using Gin + JWT + Environment Config.

📅 Week 7 — Database Integration

✅ Topics:

SQL Databases: PostgreSQL or MySQL

database/sql, connection pool

GORM ORM: models, migrations, relations

Query building, joins

Redis basics (go-redis)

Caching

💡 Mini Project:
➡️ “E-commerce Product API”
➡️ CRUD + Pagination + Redis Caching.

🗓️ PHASE 3: Advanced Backend & Open-Source Level (Week 8–11)

🎯 Goal: Production-level Go backend skills.

📅 Week 8 — Clean Architecture + Dependency Injection

✅ Topics:

Layered architecture (handlers, usecase, repository)

Dependency injection (manual or using Wire/Fx)

DTOs and validation

Modular services

💡 Mini Project:
➡️ Refactor your E-commerce API into clean architecture format.

📅 Week 9 — gRPC, REST vs gRPC

✅ Topics:

Protocol Buffers (.proto)

Generating Go code from .proto

Implementing a gRPC server and client

Comparison: REST vs gRPC

Reflection and health checks

💡 Mini Project:
➡️ Convert one of your APIs (e.g., Product API) into gRPC version.

📅 Week 10 — Testing & CI/CD

✅ Topics:

Unit testing (testing package)

Table-driven tests

Mocking interfaces (gomock)

Benchmark testing

go vet, golangci-lint

GitHub Actions for CI

💡 Mini Project:
➡️ Add tests to all your APIs (handlers + repository)
➡️ Create GitHub Actions workflow to run go test ./...

📅 Week 11 — Docker + Deployment

✅ Topics:

Dockerizing Go apps

Docker Compose (Go + PostgreSQL + Redis)

Environment configs for production

Graceful shutdown (context)

Basic Kubernetes intro

💡 Mini Project:
➡️ Dockerize your E-commerce API
➡️ Deploy to Render / Railway / AWS / Fly.io

🗓️ PHASE 4: Open Source & Job Preparation (Week 12)

🎯 Goal: Contribute to open source and prepare for job/internship.

📅 Week 12 — Open Source + Portfolio Building

✅ Topics:

Learn Git and GitHub workflow (PRs, branches, forks)

Code reviews & contribution guidelines

Reading large Go repos (Gin, Fiber, etc.)

Writing GoDocs and READMEs

Following Effective Go standards

💡 Mini Project:
➡️ Contribute small PRs to open Go projects:

Fix docs, typos, small bugs

Add test cases, examples

Improve README or refactor code

✅ Build a public portfolio with:

3–4 Go backend projects (with README + deployment link)

GitHub pinned repos

Resume mentioning contributions

🎯 FINAL OUTPUT AFTER 90 DAYS

By the end of this roadmap, you’ll be able to:
✅ Build & deploy full-scale backend APIs in Go
✅ Write clean, idiomatic, production-grade code
✅ Contribute confidently to open-source Go projects
✅ Crack backend developer interviews
✅ Stand out with a strong GitHub portfolio

🚀 Bonus (Optional Add-ons if You Have Extra Time)
Area	Why	Tool
GraphQL in Go	Modern APIs	gqlgen
Message Queues	Async jobs	RabbitMQ, Kafka
Monitoring	Observability	Prometheus, Grafana
CLI Tools	Open-source friendly	Cobra
Auth Systems	Secure APIs	OAuth2, JWT
Microservices	Distributed architecture	gRPC, Docker, NATS