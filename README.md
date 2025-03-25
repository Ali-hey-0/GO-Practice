# RSS Aggregator

<div align="center">

<img src="https://i.imgur.com/OXrGXlJ.png" alt="RSS Aggregator Logo" width="400" height="auto">

<h3>A modern RSS feed aggregator API built with Go</h3>

<p><em>Collect, manage, and consume RSS feeds through a powerful RESTful API</em></p>

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791?style=for-the-badge&logo=postgresql)](https://postgresql.org)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)](https://docker.com)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

[Features](#features) • 
[Quick Start](#getting-started) • 
[Usage](#usage-examples) • 
[API Docs](#api-documentation) • 
[Deployment](#deployment)

</div>

---

## ✨ Overview

**RSS Aggregator** is a powerful backend service that enables users to aggregate content from various RSS feeds into a single, personalized stream. Built with Go and PostgreSQL, it offers a modern, RESTful API for managing feeds and consuming content.

<div align="center">
<table>
<tr>
<td width="50%">

### Key Capabilities

- 🔐 **Secure Authentication**
- 📰 **Feed Management**
- 🔄 **Automated Content Collection**
- 🌐 **RESTful API**
- 🐳 **Docker Ready**
- 🚀 **Production Ready**

</td>
<td width="50%">

### How It Works

1. Register for an API key
2. Add your favorite RSS feeds
3. Follow feeds you're interested in
4. Our system automatically collects new posts
5. Access all your content in one place

</td>
</tr>
</table>
</div>

## 📋 Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation with Docker](#installation-with-docker)
  - [Local Development Setup](#local-development-setup)
- [Usage Examples](#usage-examples)
  - [Creating a User](#creating-a-user)
  - [Managing Feeds](#managing-feeds)
  - [Following Feeds](#following-feeds)
  - [Retrieving Posts](#retrieving-posts)
- [API Documentation](#api-documentation)
  - [Authentication](#authentication)
  - [Endpoint Reference](#endpoint-reference)
- [Database Schema](#database-schema)
- [Architecture](#architecture)
- [Docker Configuration](#docker-configuration)
- [Deployment](#deployment)
- [Performance Optimization](#performance-optimization)
- [Security Considerations](#security-considerations)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [FAQ](#faq)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## ✅ Features

<div align="center">
<table>
<tr>
<td width="50%">

### User Management 

- Create users with secure API keys
- Authenticate requests via API key validation
- Retrieve user information

### Feed Management

- Add new RSS feeds with validation
- Retrieve all available feeds
- Automatic fetching of feed metadata

</td>
<td width="50%">

### Feed Following

- Create relationships between users and feeds
- List all feeds a user is following
- Remove follows to unsubscribe from content

### Post Collection

- Automated retrieval of posts from feeds
- Deduplication to prevent duplicates
- Proper parsing of post metadata

</td>
</tr>
</table>
</div>

### Additional Features

- **Scalable Architecture**:
  - Concurrent feed processing for efficient updates
  - Background worker system separate from the API server
  - Configurable processing settings

- **RESTful API**:
  - Clean and consistent endpoint design
  - Proper status codes and error handling
  - JSON responses with clear structure

- **PostgreSQL Database**:
  - Persistent storage with efficient queries
  - Type-safe database access with SQLC
  - Proper relationship modeling

## 🛠️ Technology Stack

<div align="center">
<table>
<tr>
<td>

### Backend
- **Go 1.24+** - Modern, efficient language
- **Chi Router** - Lightweight HTTP router
- **SQLC** - Type-safe SQL query generator

</td>
<td>

### Database
- **PostgreSQL 15** - Robust relational database
- **SQL migrations** - Version-controlled schema

</td>
</tr>
<tr>
<td>

### Authentication
- **API Key-based authentication**
- Secure header-based validation

</td>
<td>

### Containerization
- **Docker** - Application containerization
- **Docker Compose** - Multi-container setup

</td>
</tr>
</table>
</div>

## 🗂️ Project Structure

```
rssagg/
├── 🐳 Dockerfile               # Docker configuration
├── 🐙 docker-compose.yml       # Docker Compose configuration
├── 📝 .env.example             # Example environment variables
├── 📦 go.mod / go.sum          # Go module definition
├── 🚀 main.go                  # Application entry point
├── 📄 json.go                  # JSON response handling
├── 🧩 models.go                # Data models and converters
├── 📁 handlers/                # API endpoint handlers
├── 🔐 middleware/              # HTTP middleware components
├── 📦 internal/                # Internal package code
│   ├── 💾 database/            # Database access
│   └── 🔑 auth/                # Authentication utilities
├── 📊 sql/                     # SQL definitions
│   ├── 🏗️ schema/             # Database schema
│   └── 🔍 queries/             # SQL queries
└── 🔄 scraper.go               # Feed scraper implementation
```

## 📋 Prerequisites

- **Docker and Docker Compose**: For containerized deployment
- **Git**: For cloning the repository
- **Go 1.24+**: For local development (not needed with Docker)
- **PostgreSQL**: For local development (not needed with Docker)

## 🚀 Getting Started

### Installation with Docker

<div align="center">

![Docker Setup](https://i.imgur.com/qnlKVcJ.png)

</div>

The quickest way to get started is using Docker Compose:

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd rssagg
   ```

2. **Start the application**:
   ```bash
   docker-compose up -d
   ```

3. **Verify the installation**:
   ```bash
   curl http://localhost:5050/v1/healthz
   # Response: {"status":"ok"}
   ```

Everything is now set up and ready to use! The API is available at `http://localhost:5050`.

### Local Development Setup

For local development without Docker:

<details>
<summary><b>Click to expand local setup instructions</b></summary>

1. **Install Go 1.24 or later**:
   Follow instructions at [golang.org](https://golang.org/doc/install)

2. **Install PostgreSQL**:
   Follow instructions at [postgresql.org](https://www.postgresql.org/download/)

3. **Create a database**:
   ```bash
   createdb postgres
   ```

4. **Set up environment variables**:
   ```bash
   cp .env.example .env
   # Edit .env file with your database credentials
   ```

5. **Run database migrations**:
   ```bash
   go run ./sql/schema
   ```

6. **Install dependencies**:
   ```bash
   go mod download
   ```

7. **Start the application**:
   ```bash
   go run .
   ```
</details>

## 💻 Usage Examples

Here are some practical examples of interacting with the API:

### Creating a User

```bash
# Create a new user
curl -X POST http://localhost:5050/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe"}'

# Response will include your API key
# Save it for future authenticated requests
API_KEY="your-api-key-from-response"
```

<details>
<summary><b>View example response</b></summary>

```json
{
  "id": "5f8d0d55-afe5-4084-b6a7-80b8a4d824b6",
  "created_at": "2023-04-15T12:34:56.789Z",
  "updated_at": "2023-04-15T12:34:56.789Z",
  "name": "John Doe",
  "api_key": "8c31059c-8527-4c1e-b810-b3a1e8590f67"
}
```
</details>

### Managing Feeds

```bash
# Add a new feed (requires authentication)
curl -X POST http://localhost:5050/v1/feeds \
  -H "Content-Type: application/json" \
  -H "Authorization: ApiKey $API_KEY" \
  -d '{"name": "Go Blog", "url": "https://go.dev/blog/feed.atom"}'
```

<details>
<summary><b>See more feed operations</b></summary>

```bash
# Get all feeds
curl http://localhost:5050/v1/feeds
```
</details>

### Following Feeds

```bash
# Follow a feed (requires authentication)
curl -X POST http://localhost:5050/v1/feed_follows \
  -H "Content-Type: application/json" \
  -H "Authorization: ApiKey $API_KEY" \
  -d '{"feed_id": "feed-uuid-from-previous-response"}'
```

<details>
<summary><b>See more feed follow operations</b></summary>

```bash
# Get your followed feeds
curl -H "Authorization: ApiKey $API_KEY" \
  http://localhost:5050/v1/feed_follows

# Unfollow a feed
curl -X DELETE -H "Authorization: ApiKey $API_KEY" \
  http://localhost:5050/v1/feed_follows/feed-follow-id
```
</details>

### Retrieving Posts

```bash
# Get posts from followed feeds with a limit
curl -H "Authorization: ApiKey $API_KEY" \
  http://localhost:5050/v1/posts?limit=20
```

## 📖 API Documentation

### Authentication

Most endpoints require an API key provided in the `Authorization` header:

```
Authorization: ApiKey YOUR_API_KEY
```

<details>
<summary><b>View Authentication Details</b></summary>

You can obtain an API key by creating a user through the `/v1/users` endpoint.

All authenticated endpoints will return:
- `401 Unauthorized` if no API key is provided
- `401 Unauthorized` if an invalid API key is provided
</details>

### Endpoint Reference

<div align="center">
<table>
<tr>
<th>Endpoint</th>
<th>Method</th>
<th>Description</th>
<th>Auth Required</th>
</tr>
<tr>
<td><code>/v1/healthz</code></td>
<td>GET</td>
<td>Health check</td>
<td>No</td>
</tr>
<tr>
<td><code>/v1/users</code></td>
<td>POST</td>
<td>Create a new user</td>
<td>No</td>
</tr>
<tr>
<td><code>/v1/users</code></td>
<td>GET</td>
<td>Get authenticated user</td>
<td>Yes</td>
</tr>
<tr>
<td><code>/v1/feeds</code></td>
<td>POST</td>
<td>Create a new feed</td>
<td>Yes</td>
</tr>
<tr>
<td><code>/v1/feeds</code></td>
<td>GET</td>
<td>Get all feeds</td>
<td>No</td>
</tr>
<tr>
<td><code>/v1/feed_follows</code></td>
<td>POST</td>
<td>Follow a feed</td>
<td>Yes</td>
</tr>
<tr>
<td><code>/v1/feed_follows</code></td>
<td>GET</td>
<td>Get user's feed follows</td>
<td>Yes</td>
</tr>
<tr>
<td><code>/v1/feed_follows/{id}</code></td>
<td>DELETE</td>
<td>Unfollow a feed</td>
<td>Yes</td>
</tr>
<tr>
<td><code>/v1/posts</code></td>
<td>GET</td>
<td>Get posts from followed feeds</td>
<td>Yes</td>
</tr>
</table>
</div>

<details>
<summary><b>View detailed API documentation</b></summary>

#### Health Check

**Endpoint**: `GET /v1/healthz`

**Description**: Check if the server is running

**Response**:
```json
{
  "status": "ok"
}
```

**Status Codes**:
- `200 OK`: Server is running

#### Users

**Endpoint**: `POST /v1/users`

**Description**: Create a new user

**Request Body**:
```json
{
  "name": "Your Name"
}
```

**Response**:
```json
{
  "id": "5f8d0d55-afe5-4084-b6a7-80b8a4d824b6",
  "created_at": "2023-04-15T12:34:56.789Z",
  "updated_at": "2023-04-15T12:34:56.789Z",
  "name": "Your Name",
  "api_key": "8c31059c-8527-4c1e-b810-b3a1e8590f67"
}
```

**Status Codes**:
- `201 Created`: User successfully created
- `400 Bad Request`: Invalid input

---

**Endpoint**: `GET /v1/users`

**Description**: Get authenticated user

**Authentication**: Required

**Response**:
```json
{
  "id": "5f8d0d55-afe5-4084-b6a7-80b8a4d824b6",
  "created_at": "2023-04-15T12:34:56.789Z",
  "updated_at": "2023-04-15T12:34:56.789Z",
  "name": "Your Name"
}
```

**Status Codes**:
- `200 OK`: User information retrieved
- `401 Unauthorized`: Missing or invalid API key

#### Feeds

**Endpoint**: `POST /v1/feeds`

**Description**: Create a new feed

**Authentication**: Required

**Request Body**:
```json
{
  "name": "Feed Name",
  "url": "https://example.com/feed.xml"
}
```

**Response**:
```json
{
  "id": "7f8c0c44-afe5-4084-b6a7-80b8a4d824b6",
  "created_at": "2023-04-15T12:34:56.789Z",
  "updated_at": "2023-04-15T12:34:56.789Z",
  "name": "Feed Name",
  "url": "https://example.com/feed.xml",
  "user_id": "5f8d0d55-afe5-4084-b6a7-80b8a4d824b6"
}
```

**Status Codes**:
- `201 Created`: Feed successfully created
- `400 Bad Request`: Invalid input
- `401 Unauthorized`: Missing or invalid API key

---

**Endpoint**: `GET /v1/feeds`

**Description**: Get all feeds

**Response**: Array of feed objects

**Status Codes**:
- `200 OK`: Feeds retrieved successfully

#### Feed Follows

**Endpoint**: `POST /v1/feed_follows`

**Description**: Follow a feed

**Authentication**: Required

**Request Body**:
```json
{
  "feed_id": "7f8c0c44-afe5-4084-b6a7-80b8a4d824b6"
}
```

**Response**: Feed follow object

**Status Codes**:
- `201 Created`: Feed follow created successfully
- `400 Bad Request`: Invalid input
- `401 Unauthorized`: Missing or invalid API key
- `404 Not Found`: Feed not found

---

**Endpoint**: `GET /v1/feed_follows`

**Description**: Get all feed follows for a user

**Authentication**: Required

**Response**: Array of feed follow objects

**Status Codes**:
- `200 OK`: Feed follows retrieved successfully
- `401 Unauthorized`: Missing or invalid API key

---

**Endpoint**: `DELETE /v1/feed_follows/{feedFollowID}`

**Description**: Unfollow a feed

**Authentication**: Required

**Status Codes**:
- `204 No Content`: Feed follow deleted successfully
- `401 Unauthorized`: Missing or invalid API key
- `404 Not Found`: Feed follow not found

#### Posts

**Endpoint**: `GET /v1/posts`

**Description**: Get posts for authenticated user

**Authentication**: Required

**Query Parameters**:
- `limit`: Maximum number of posts to return (default: 10)

**Response**: Array of post objects

**Status Codes**:
- `200 OK`: Posts retrieved successfully
- `401 Unauthorized`: Missing or invalid API key
</details>

## 💾 Database Schema

<div align="center">

![Database Schema](https://i.imgur.com/fsJeXJJ.png)

</div>

<details>
<summary><b>View detailed database schema</b></summary>

### `users` Table
| Column | Type | Description |
|--------|------|-------------|
| id | UUID | Primary key |
| created_at | TIMESTAMP | Creation timestamp |
| updated_at | TIMESTAMP | Update timestamp |
| name | TEXT | User's name |
| api_key | TEXT | Authentication key |

### `feeds` Table
| Column | Type | Description |
|--------|------|-------------|
| id | UUID | Primary key |
| created_at | TIMESTAMP | Creation timestamp |
| updated_at | TIMESTAMP | Update timestamp |
| name | TEXT | Feed name |
| url | TEXT | Feed URL (unique) |
| user_id | UUID | Creator reference (FK) |
| last_fetched_at | TIMESTAMP | Last fetch time |

### `feed_follows` Table
| Column | Type | Description |
|--------|------|-------------|
| id | UUID | Primary key |
| created_at | TIMESTAMP | Creation timestamp |
| updated_at | TIMESTAMP | Update timestamp |
| user_id | UUID | User reference (FK) |
| feed_id | UUID | Feed reference (FK) |

### `posts` Table
| Column | Type | Description |
|--------|------|-------------|
| id | UUID | Primary key |
| created_at | TIMESTAMP | Creation timestamp |
| updated_at | TIMESTAMP | Update timestamp |
| title | TEXT | Post title |
| url | TEXT | Post URL (unique) |
| description | TEXT | Post content |
| published_at | TIMESTAMP | Publication date |
| feed_id | UUID | Feed reference (FK) |
</details>

## 🏛️ Architecture

<div align="center">

![Architecture Diagram](https://i.imgur.com/lnXPLn1.png)

</div>

<details>
<summary><b>Architecture Details</b></summary>

The RSS Aggregator follows a clean architecture pattern:

1. **HTTP Server (Chi Router)**: 
   - Handles API requests and routing
   - Includes middleware for auth, CORS, etc.

2. **Business Logic (Handlers)**:
   - Contains the application logic
   - Validates inputs and prepares responses

3. **Database Layer (SQLC)**:
   - Type-safe database interactions
   - Converts between models

4. **Feed Scraper (Background Worker)**:
   - Runs on a schedule
   - Fetches and processes RSS feeds
   - Stores posts in the database
</details>

## 🐳 Docker Configuration

The application is containerized using Docker:

<details>
<summary><b>View Docker configuration</b></summary>

### `Dockerfile`

A multi-stage build process for the Go application:
```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rssagg .

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/rssagg .
EXPOSE 5050
CMD ["./rssagg"]
```

### `docker-compose.yml`

Orchestrates the application and database:
```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "5050:5050"
    depends_on:
      - db
    environment:
      - PORT=5050
      - DB_URL=postgres://postgres:2000@db:5432/postgres?sslmode=disable
    restart: always

  db:
    image: postgres:15-alpine
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=2000
      - POSTGRES_DB=postgres
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./sql/schema:/docker-entrypoint-initdb.d
    restart: always

volumes:
  postgres_data:
```
</details>

## 🚀 Deployment

<div align="center">

![Deployment](https://i.imgur.com/9MJRbBV.png)

</div>

<details>
<summary><b>Production Deployment Guidelines</b></summary>

For production deployment, consider:

1. **Database Management**:
   - Use a managed PostgreSQL service
   - Implement proper backup procedures
   - Set up database replicas

2. **Security**:
   - Implement proper SSL/TLS
   - Use a reverse proxy
   - Implement rate limiting
   - Use a secrets manager

3. **Monitoring and Logging**:
   - Set up centralized logging
   - Implement application metrics
   - Configure alerting

4. **Scaling**:
   - Deploy multiple instances
   - Consider Kubernetes for orchestration
   - Separate API and scraper services
</details>

## ⚡ Performance Optimization

<details>
<summary><b>Performance Tips</b></summary>

1. **Database Optimization**:
   - Add indexes to frequently queried columns
   - Consider table partitioning for large datasets
   - Implement query caching

2. **Feed Scraper Performance**:
   - Use exponential backoff for retries
   - Optimize scraping intervals
   - Consider a queue system for processing

3. **API Response Time**:
   - Implement pagination
   - Consider Redis for caching
   - Optimize database queries
</details>

## 🔒 Security Considerations

<details>
<summary><b>Security Best Practices</b></summary>

1. **API Key Management**:
   - Store hashed keys
   - Implement key rotation
   - Consider scoped keys

2. **Input Validation**:
   - Validate all inputs
   - Sanitize content
   - Implement proper error handling

3. **Content Security**:
   - Sanitize feed content
   - Validate content before storage
   - Set content size limits
</details>

## ⚙️ Configuration

<details>
<summary><b>Environment Variables</b></summary>

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| PORT | HTTP server port | 5050 | 8080 |
| DB_URL | PostgreSQL connection string | - | postgres://user:pass@host:port/dbname |
| SCRAPER_CONCURRENCY | Concurrent feed scrapers | 10 | 20 |
| SCRAPER_INTERVAL | Scraping interval (minutes) | 60 | 30 |
| LOG_LEVEL | Logging verbosity | info | debug |

Example `.env` file:
```
PORT=5050
DB_URL=postgresql://postgres:2000@127.0.0.1:5432/postgres?sslmode=disable
SCRAPER_CONCURRENCY=10
SCRAPER_INTERVAL=60
LOG_LEVEL=info
```
</details>

## 🔍 Troubleshooting

<details>
<summary><b>Common Issues and Solutions</b></summary>

### Application Won't Start

**Check database connection:**
```bash
# Verify PostgreSQL is running
pg_isready -h localhost -p 5432

# Check connection string
echo $DB_URL
```

### Authentication Failures

**Verify API key format:**
```
Authorization: ApiKey YOUR_API_KEY
```

### No Posts Appearing

**Ensure you have followed feeds:**
```bash
curl -H "Authorization: ApiKey $API_KEY" \
  http://localhost:5050/v1/feed_follows
```
</details>

## ❓ FAQ

<details>
<summary><b>Frequently Asked Questions</b></summary>

**Q: How frequently are feeds checked?**  
A: By default, every 60 minutes. Configurable via `SCRAPER_INTERVAL`.

**Q: Can I follow feeds created by others?**  
A: Yes, the `/v1/feeds` endpoint returns all feeds.

**Q: Is there a limit to feeds I can follow?**  
A: No built-in limit, but consider performance with large numbers.
</details>

## 👥 Contributing

<details>
<summary><b>Contribution Guidelines</b></summary>

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

Please ensure your code follows the existing style and includes tests.
</details>

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- The Go community for excellent libraries and tools
- PostgreSQL for a robust database solution
- The RSS standard for enabling content syndication
- Open source contributors everywhere