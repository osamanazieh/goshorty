# Shorty

A simple URL shortener REST API built with **Go** and **PostgreSQL**.

Project URL: https://github.com/osamanazieh/goshorty

Shorty allows you to create shortened URLs, retrieve the original URLs using their short codes, update and delete existing URLs, and track how many times a shortened URL has been accessed.

## Features

* Generate short, random URL codes
* Store shortened URLs in PostgreSQL
* Retrieve the original URL using a short code
* Track URL access count
* Update existing URLs
* Delete shortened URLs
* RESTful HTTP API
* JSON request/response handling
* Database queries generated with `sqlc`
* PostgreSQL database integration

## Tech Stack

* **Go**
* **net/http**
* **PostgreSQL**
* **sqlc**
* **godotenv**
* **github.com/google/uuid**
* **github.com/lib/pq**

## Project Structure

```text
shorty/
├── cmd/
│   └── ...
├── internal/
│   ├── API/
│   │   ├── configuration.go
│   │   ├── handlers.go
│   │   ├── generateUrl.go
│   │   ├── retrieveOriginalURL.go
│   │   ├── updateURL.go
│   │   ├── deleteUrl.go
│   │   └── getStats.go
│   │
│   └── database/
│       └── ...
│
├── migrations/
│   └── ...
│
├── .env
├── go.mod
└── main.go
```

The API configuration contains the generated database query interface used by the handlers.

## Getting Started

### Prerequisites

Make sure you have the following installed:

* Go
* PostgreSQL
* sqlc
* Git

### Clone the repository

```bash
git clone <repository-url>
cd shorty
```

### Environment Variables

Create a `.env` file in the project root:

```env
DB_URL=postgres://username:password@localhost:5432/shorty?sslmode=disable
```

The application loads the database URL from the `DB_URL` environment variable and uses it to establish the PostgreSQL connection.

### Database

Create a PostgreSQL database named `shorty`:

```sql
CREATE DATABASE shorty;
```

Run the project's database migrations before starting the server.

## Running the Server

Start the API with:

```bash
go run .
```

The server listens on:

```text
http://localhost:8080
```

The current server registers the API endpoints directly on a `http.ServeMux`.

---

# API Reference

## Create a Short URL

Creates a new shortened URL.

### Request

```http
POST /shorten
Content-Type: application/json
```

```json
{
  "url": "https://example.com/some/very/long/url"
}
```

### Response

**201 Created**

```json
{
  "id": "...",
  "url": "https://example.com/some/very/long/url",
  "short_code": "..."
}
```

## The server generates 8 random bytes and encodes them using Base64 to create the short code. A UUID is also generated for the database record, along with creation and update timestamps.

## Retrieve Original URL

Retrieves the URL associated with a short code.

### Request

```http
GET /shorten/{short_code}
```

Example:

```http
GET /shorten/AbCdEfGh
```

Every successful retrieval increments the URL's hit counter.

### Response

**200 OK**

```json
{
  "id": "...",
  "url": "https://example.com",
  "short_code": "AbCdEfGh",
  "hits": 1
}
```

---

## Update a Short URL

Updates the original URL associated with a short code.

### Request

```http
PUT /shorten/{short_code}
Content-Type: application/json
```

```json
{
  "url": "https://example.com/new-url"
}
```

### Response

**200 OK**

The URL and `updated_at` timestamp are updated in the database.

---

## Delete a Short URL

Deletes a shortened URL.

### Request

```http
DELETE /shorten/{short_code}
```

Example:

```http
DELETE /shorten/AbCdEfGh
```

### Response

**204 No Content**

The handler deletes the database record associated with the supplied short code.

---

## Get URL Statistics

Retrieves information about a shortened URL.

### Request

```http
DELETE /shorten/{short_code}/stats
```

> **Note:** The current implementation registers this endpoint as `DELETE`, even though the handler only retrieves statistics. If this endpoint is intended to be read-only, `GET` would be a more appropriate HTTP method.

### Response

**200 OK**

The endpoint retrieves the URL record associated with the short code.

---

# Example Workflow

### 1. Create a shortened URL

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'
```

### 2. Retrieve it

```bash
curl http://localhost:8080/shorten/AbCdEfGh
```

Each successful retrieval increases the URL's hit count.

### 3. Update it

```bash
curl -X PUT http://localhost:8080/shorten/AbCdEfGh \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/new"}'
```

### 4. Delete it

```bash
curl -X DELETE http://localhost:8080/shorten/AbCdEfGh
```

---

# Database

Shorty uses PostgreSQL as its persistent storage layer.

Database access is abstracted through generated queries. The API configuration receives a `*database.Queries` instance, allowing HTTP handlers to interact with the database without directly managing SQL connections.

The application initializes the database with:

```go
dbQueries := database.New(db)
```

and injects it into the API configuration.

---

# Design

The application follows a relatively simple layered structure:

```text
HTTP Request
     │
     ▼
HTTP Handler
     │
     ▼
API Configuration
     │
     ▼
sqlc Generated Queries
     │
     ▼
PostgreSQL
```

Handlers extract path parameters using Go's `Request.PathValue`, validate JSON request bodies, execute database operations, and return JSON responses.

Common JSON handling is centralized in helper functions such as `respondWithJSON` and the generic `getJSON` function.

---

# Future Improvements

Possible improvements for the project include:

* [ ] Redirect directly to the original URL instead of returning it as JSON
* [ ] Add URL expiration
* [ ] Add custom short codes
* [ ] Validate submitted URLs
* [ ] Prevent short-code collisions
* [ ] Add authentication and authorization
* [ ] Add rate limiting
* [ ] Add request logging
* [ ] Add structured error responses
* [ ] Add automated tests
* [ ] Add integration tests
* [ ] Add Docker support
* [ ] Add API documentation with OpenAPI/Swagger
* [ ] Add proper HTTP redirects for shortened URLs

# License

This project is for educational and personal development purposes.
